///
//  Generated code. Do not modify.
//  source: payeet.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'payeet.pb.dart' as $0;
export 'payeet.pb.dart';

class payeetClient extends $grpc.Client {
  static final _$login = $grpc.ClientMethod<$0.LoginRequest, $0.LoginResponse>(
      '/payeet.payeet/Login',
      ($0.LoginRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.LoginResponse.fromBuffer(value));
  static final _$loginS =
      $grpc.ClientMethod<$0.LoginRequest_S, $0.LoginResponse>(
          '/payeet.payeet/LoginS',
          ($0.LoginRequest_S value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.LoginResponse.fromBuffer(value));
  static final _$register =
      $grpc.ClientMethod<$0.RegisterRequest, $0.StatusResponse>(
          '/payeet.payeet/Register',
          ($0.RegisterRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.StatusResponse.fromBuffer(value));
  static final _$getBalance =
      $grpc.ClientMethod<$0.BalanceRequest, $0.BalanceResponse>(
          '/payeet.payeet/GetBalance',
          ($0.BalanceRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.BalanceResponse.fromBuffer(value));
  static final _$transferBalance =
      $grpc.ClientMethod<$0.TransferRequest, $0.StatusResponse>(
          '/payeet.payeet/TransferBalance',
          ($0.TransferRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.StatusResponse.fromBuffer(value));

  payeetClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options,
      $core.Iterable<$grpc.ClientInterceptor> interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.LoginResponse> login($0.LoginRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$login, request, options: options);
  }

  $grpc.ResponseFuture<$0.LoginResponse> loginS($0.LoginRequest_S request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$loginS, request, options: options);
  }

  $grpc.ResponseFuture<$0.StatusResponse> register($0.RegisterRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$register, request, options: options);
  }

  $grpc.ResponseFuture<$0.BalanceResponse> getBalance($0.BalanceRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$getBalance, request, options: options);
  }

  $grpc.ResponseFuture<$0.StatusResponse> transferBalance(
      $0.TransferRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$transferBalance, request, options: options);
  }
}

abstract class payeetServiceBase extends $grpc.Service {
  $core.String get $name => 'payeet.payeet';

  payeetServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.LoginRequest, $0.LoginResponse>(
        'Login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.LoginRequest.fromBuffer(value),
        ($0.LoginResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.LoginRequest_S, $0.LoginResponse>(
        'LoginS',
        loginS_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.LoginRequest_S.fromBuffer(value),
        ($0.LoginResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RegisterRequest, $0.StatusResponse>(
        'Register',
        register_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RegisterRequest.fromBuffer(value),
        ($0.StatusResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.BalanceRequest, $0.BalanceResponse>(
        'GetBalance',
        getBalance_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.BalanceRequest.fromBuffer(value),
        ($0.BalanceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.TransferRequest, $0.StatusResponse>(
        'TransferBalance',
        transferBalance_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.TransferRequest.fromBuffer(value),
        ($0.StatusResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.LoginResponse> login_Pre(
      $grpc.ServiceCall call, $async.Future<$0.LoginRequest> request) async {
    return login(call, await request);
  }

  $async.Future<$0.LoginResponse> loginS_Pre(
      $grpc.ServiceCall call, $async.Future<$0.LoginRequest_S> request) async {
    return loginS(call, await request);
  }

  $async.Future<$0.StatusResponse> register_Pre(
      $grpc.ServiceCall call, $async.Future<$0.RegisterRequest> request) async {
    return register(call, await request);
  }

  $async.Future<$0.BalanceResponse> getBalance_Pre(
      $grpc.ServiceCall call, $async.Future<$0.BalanceRequest> request) async {
    return getBalance(call, await request);
  }

  $async.Future<$0.StatusResponse> transferBalance_Pre(
      $grpc.ServiceCall call, $async.Future<$0.TransferRequest> request) async {
    return transferBalance(call, await request);
  }

  $async.Future<$0.LoginResponse> login(
      $grpc.ServiceCall call, $0.LoginRequest request);
  $async.Future<$0.LoginResponse> loginS(
      $grpc.ServiceCall call, $0.LoginRequest_S request);
  $async.Future<$0.StatusResponse> register(
      $grpc.ServiceCall call, $0.RegisterRequest request);
  $async.Future<$0.BalanceResponse> getBalance(
      $grpc.ServiceCall call, $0.BalanceRequest request);
  $async.Future<$0.StatusResponse> transferBalance(
      $grpc.ServiceCall call, $0.TransferRequest request);
}
