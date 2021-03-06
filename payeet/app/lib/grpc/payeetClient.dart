import 'dart:io';

import 'package:Payeet/Screens/VerifyPage.dart';
import 'package:fixnum/fixnum.dart';
import 'package:device_info/device_info.dart';
import 'package:flutter/services.dart';

import 'package:grpc/grpc.dart';
import 'package:grpc/service_api.dart' as $grpc;

import 'dart:async';

import 'protos/payeet.pb.dart';
import 'protos/payeet.pbgrpc.dart';

import '../secure_storage.dart';

class PayeetClient {
  final SecureStorage secureStorage = SecureStorage();
  final PayeetChannel channel;

  String _accessToken;
  Int64 tokenExpiresOn;
  String _refreshToken;

  payeet_authClient _unauthenticatedClient;
  payeetClient _authenticatedClient;

  bool
      _cachedBalance; // this is true when the client class has cached the user's balance
  String _balance;

  bool
      _cachedInfo; // this is true when the client class has cached the user info from the server
  String _firstName;
  String _lastName;
  List<String> _friends;
  List<String> _followers;

  List<UserInfoResponse> _topUsers;

  String _userID;

  // ctor
  PayeetClient(this.channel) {
    _cachedInfo = false;
    _unauthenticatedClient = payeet_authClient(channel);
  }

  bool get cachedBalance => _cachedBalance;
  String get getCachedBalance => _balance;

  bool get cachedInfo => _cachedInfo;
  String get getCachedFirstName => _firstName;
  String get getCachedLastName => _lastName;
  List<String> get getCachedFriends => _friends;
  List<String> get getCachedFollowers => _followers;
  List<UserInfoResponse> get getTopUsers => _topUsers;
  String get getCachedUserID => _userID;

  Future<LoginResponse> login(String mail, String password) async {
    LoginResponse response;

    String identifier = '';
    final DeviceInfoPlugin deviceInfoPlugin = new DeviceInfoPlugin();
    try {
      if (Platform.isAndroid) {
        var build = await deviceInfoPlugin.androidInfo;
        identifier = build.androidId;  //UUID for Android
      } else if (Platform.isIOS) {
        var data = await deviceInfoPlugin.iosInfo;
        identifier = data.identifierForVendor;  //UUID for iOS
      }
    } on PlatformException {
      print('Failed to get platform version');
    }

    print(identifier);
    try {
      response = await _unauthenticatedClient.login(LoginRequest()
        ..mail = mail
        ..password = password
        ..identifier = identifier);
    } catch (e) {
      rethrow; // cant login so throw the error
    }

    _accessToken = response.accessToken;
    tokenExpiresOn = response.expiresOn;
    _refreshToken = response.refreshToken;

    await SecureStorage.writeSecureData('refreshToken', _refreshToken);

    channel.setAccessTokenMetadata(_accessToken);
    _authenticatedClient =
        payeetClient(channel); // create the authenticated client

    await getUserInfo();

    return response;
  }

  Future<bool> loginWithRefresh() async {
    _refreshToken = await SecureStorage.readSecureData('refreshToken');
    LoginResponse response;

    try {
      response = await _unauthenticatedClient
          .refreshToken(RefreshTokenRequest()..refreshToken = _refreshToken);
    } catch (e) {
      //rethrow; // cant login so throw the error
      return false;
    }

    _accessToken = response.accessToken;
    tokenExpiresOn = response.expiresOn;
    _refreshToken = response.refreshToken;

    await SecureStorage.writeSecureData('refreshToken', _refreshToken);

    channel.setAccessTokenMetadata(_accessToken);
    _authenticatedClient =
        payeetClient(channel); // create the authenticated client

    await getUserInfo();

    return true;
  }

  Future<StatusResponse> register(
      String firstName, String lastName, String mail, String password) async {
    final response = await _unauthenticatedClient.register(RegisterRequest()
          ..firstName = firstName
          ..lastName = lastName
          ..mail = mail
          ..password =
              password // add a small hash logic as sending a plain text password isnt a good practice
        );

    return response;
  }

  ResponseStream<HistoryResponse> getTransferHistory(String mail) {
    final response = _authenticatedClient
        .getFullHistory(HistoryRequest()..senderMail = mail);

    return response;
  }

  ResponseStream<SearchFriendResponse> searchFriend(String text) {
    final response =
        _authenticatedClient.searchFriend(SearchFriendRequest()..search = text);

    return response;
  }

  Future<LoginResponse> refreshAccessToken() async {
    final response = await _unauthenticatedClient
        .refreshToken(RefreshTokenRequest()..refreshToken = _refreshToken);

    _accessToken = response.accessToken;
    tokenExpiresOn = response.expiresOn;
    _refreshToken = response.refreshToken;
    channel.setAccessTokenMetadata(_accessToken);

    return response;
  }

  Future<BalanceResponse> getBalance() async {
    final response = await _authenticatedClient.getBalance(BalanceRequest());

    _cachedBalance = true;
    _balance = response.balance;

    return response;
  }

  Future<UserInfoResponse> getUserInfo() async {
    final response = await _authenticatedClient.getUserInfo(UserInfoRequest());

    // caching the user info
    _firstName = response.firstName;
    _lastName = response.lastName;
    _userID = response.mail;

    return response;
  }

  Future<StatusResponse> transferBalance(String mail, int amount) async {
    final response =
        await _authenticatedClient.transferBalance(TransferRequest()
          ..receiverMail = mail
          ..amount = amount);

    return response;
  }

  Future<StatusResponse> addFriend(String mail) async {
    final response =
        await _authenticatedClient.addFriend(AddFriendRequest()..mail = mail);

    return response;
  }

  Future<StatusResponse> verify(String code, String mail) async {
    final response = await _unauthenticatedClient.verify(VerifyRequest()
      ..code = code
      ..mail = mail);

    return response;
  }

  Future<StatusResponse> removeFriend(String mail) async {
    final response = await _authenticatedClient
        .removeFriend(RemoveFriendRequest()..mail = mail);

    return response;
  }

  Future<void> getFriends() async {
    List<GetFriendsResponse> d =
        await _authenticatedClient.getFriends(GetFriendsRequest()).toList();
    _friends = d.map((e) => e.mail).toList();
  }

  Future<void> fetchFollowers() async {
    var d =
        await _authenticatedClient.getFollowers(GetFollowersRequest()).toList();
    _followers = d.map((e) => e.mail).toList();
  }

  Future<void> fetchTopUsers() async {
    final response = await _authenticatedClient.getTopUsers(TopUsersRequest());

    this._topUsers = response.users;
  }
}

// implementing the ClientChannel to have an interceptor to set the authorization
// metadata header when each request is invoked
class PayeetChannel implements $grpc.ClientChannel {
  final $grpc.ClientChannel channel;
  CallOptions _options;

  PayeetChannel(this.channel); // ctor

  @override
  Future<void> shutdown() => channel.shutdown();

  @override
  Future<void> terminate() => channel.terminate();

  @override
  ClientCall<Q, R> createCall<Q, R>(
      ClientMethod<Q, R> method, Stream<Q> requests, CallOptions options) {
    return channel.createCall<Q, R>(
        method, requests, options.mergedWith(_options));
  }

  void setAccessTokenMetadata(String accessToken) {
    _options = CallOptions(metadata: {'authorization': accessToken});
  }
}
