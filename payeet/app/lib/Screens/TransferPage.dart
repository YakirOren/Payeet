import 'package:flappy_search_bar/search_bar_style.dart';
import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';
import 'package:Payeet/globals.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:intl/intl.dart';
import 'package:Payeet/UI_Elements/confirm.dart';
import 'package:keyboard_actions/keyboard_actions.dart';
import 'package:flappy_search_bar/flappy_search_bar.dart';
import 'package:Payeet/grpc/protos/payeet.pbgrpc.dart';

class TransferPage extends StatefulWidget {
  @override
  _TransferPageState createState() => _TransferPageState();
}

class _TransferPageState extends State<TransferPage> {
  final emailControler = TextEditingController();

  @override
  void dispose() {
    // Clean up the controller when the widget is disposed.
    emailControler.dispose();
    super.dispose();
  }

  final _formKey = GlobalKey<FormState>();
  bool _loading = false;

  int index = 1;

  int selected_index = 0;

  bool showCubes = true;
  final _nodeText1 = FocusNode();
  final _nodeText2 = FocusNode();

  KeyboardActionsConfig _buildConfig(BuildContext context) {
    return KeyboardActionsConfig(
      keyboardActionsPlatform: KeyboardActionsPlatform.ALL,
      keyboardBarColor: Colors.grey[200],
      nextFocus: true,
      actions: [
        KeyboardActionsItem(focusNode: _nodeText1, toolbarButtons: [
          (node) {
            return GestureDetector(
              onTap: () => node.unfocus(),
              child: Padding(
                padding: EdgeInsets.all(8.0),
                child: Icon(Icons.close),
              ),
            );
          }
        ]),
        KeyboardActionsItem(
          focusNode: _nodeText2,
        ),
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    TextStyle style = TextStyle(
        fontFamily: 'Montserrat',
        fontSize: 20.0,
        color: Theme.of(context).highlightColor);

    return KeyboardActions(
      config: _buildConfig(context),
      child: Container(
          //resizeToAvoidBottomInset: false,
          child: Column(
        children: <Widget>[
          Container(
            child: Consumer(builder: (context, watch, _) {
              final balance = watch(Globals.balance).state;
              final formatCurrency =
                  new NumberFormat.simpleCurrency(decimalDigits: 0);
              return FittedBox(
                child: RichText(
                  text: TextSpan(
                    text: "Balance\n",
                    style: TextStyle(
                        fontWeight: FontWeight.w300,
                        color: Theme.of(context).highlightColor,
                        fontSize: 30),
                    children: <TextSpan>[
                      TextSpan(
                        text: "${formatCurrency.format(balance)}",
                        style: TextStyle(
                            fontWeight: FontWeight.w300,
                            color: Theme.of(context).highlightColor,
                            fontSize: 40),
                      ),
                    ],
                  ),
                ),
              );
            }),
          ),

          Column(
            children: [
              Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 36.0),
                  child: Column(
                    children: [
                      ButtonBar(
                        children: [
                          TextButton(
                              onPressed: () {
                                setState(() {
                                  index = 1;
                                });
                              },
                              child: Text(
                                "Favorites",
                                style: TextStyle(
                                    fontWeight: FontWeight.w300,
                                    color: Theme.of(context).highlightColor,
                                    fontSize: 20),
                              )),
                          TextButton(
                              onPressed: () {
                                setState(() {
                                  index = 0;
                                });
                              },
                              child: Text(
                                "Add Friend",
                                style: TextStyle(
                                    fontWeight: FontWeight.w300,
                                    color: Theme.of(context).highlightColor,
                                    fontSize: 20),
                              )),
                        ],
                      ),
                      index == 1
                          ? ButtonBar(
                              children: [
                                IconButton(
                                  onPressed: () {
                                    setState(() {
                                      showCubes = true;
                                    });
                                  },
                                  icon: Icon(
                                    Icons.view_week,
                                    color: Theme.of(context).highlightColor,
                                  ),
                                ),
                                IconButton(
                                  onPressed: () {
                                    setState(() {
                                      showCubes = false;
                                    });
                                  },
                                  icon: Icon(
                                    Icons.list,
                                    color: Theme.of(context).highlightColor,
                                  ),
                                ),
                              ],
                            )
                          : Container(),
                    ],
                  )),
              Padding(
                padding: const EdgeInsets.only(bottom: 36, left: 36, right: 36),
                child: index == 0
                    ? LimitedBox(maxHeight: 400, child: SearchFriend())
                    : LimitedBox(
                        maxHeight: 105,
                        child: showCubes
                            ? ListView.separated(
                                scrollDirection: Axis.horizontal,
                                itemBuilder: (_, index) => GestureDetector(
                                  onLongPress: () {
                                    showDialog(
                                      context: context,
                                      builder: (_) => ConfirmDialog(
                                        danger: true,
                                        title:
                                            'Delete ${Globals.client.getCachedFriends[index]}?',
                                        cancelFunction: () {
                                          setState(() {
                                            Globals.client.getCachedFriends;
                                          });
                                          Navigator.of(context).pop();
                                          return false;
                                        },
                                        actionText: Text('Approve'),
                                        actionFunction: () async {
                                          try {
                                            await Globals.client.removeFriend(
                                                Globals.client
                                                    .getCachedFriends[index]);
                                            setState(() {
                                              Globals.client.getCachedFriends
                                                  .removeAt(index);
                                            });
                                          } catch (e) {
                                            ScaffoldMessenger.of(context)
                                                .showSnackBar(SnackBar(
                                              content: Text(
                                                  '[${e.codeName}] ${e.message}'),
                                              backgroundColor: Colors.red,
                                            ));
                                          }

                                          Navigator.of(context).pop();
                                          return true;
                                        },
                                      ),
                                    );
                                  },
                                  onTap: () {
                                    setState(() {
                                      context
                                              .read(Globals.transfer_email)
                                              .state =
                                          Globals
                                              .client.getCachedFriends[index];
                                      selected_index = index;
                                    });
                                  },
                                  child: SizedBox(
                                    width: 200,
                                    child: Container(
                                      decoration: BoxDecoration(
                                          color: selected_index == index
                                              ? Colors.cyan
                                              : Theme.of(context)
                                                  .highlightColor,
                                          // border: Border.all(
                                          //   color: Colors.red[500],
                                          // ),
                                          borderRadius: BorderRadius.all(
                                              Radius.circular(20))),
                                      child: Column(
                                        mainAxisAlignment:
                                            MainAxisAlignment.center,
                                        children: [
                                          // SizedBox(
                                          //   width:
                                          //       MediaQuery.of(context).size.width / 4,
                                          // ),
                                          Container(
                                            child: CircleAvatar(
                                              radius: 30,
                                              backgroundImage: AssetImage(
                                                  'assets/images/avatar.png'),
                                            ),
                                          ),

                                          Center(
                                            child: RichText(
                                                text: TextSpan(
                                              style: Theme.of(context)
                                                  .textTheme
                                                  .headline2
                                                  .copyWith(
                                                      color: Theme.of(context)
                                                          .accentColor),
                                              text:
                                                  "${Globals.client.getCachedFriends[index]}\n",
                                            )),
                                          ),
                                        ],
                                      ),
                                    ),
                                  ),
                                ),
                                separatorBuilder: (_, index) => VerticalDivider(
                                  width: 20,
                                ),
                                itemCount:
                                    Globals.client.getCachedFriends.length,
                              )
                            : ListView.separated(
                                itemBuilder: (_, index) => Padding(
                                    padding: const EdgeInsets.only(bottom: 3),
                                    child: Dismissible(
                                        confirmDismiss: (direction) async {
                                          if (direction ==
                                              DismissDirection.endToStart) {
                                            setState(() {
                                              context
                                                  .read(Globals.transfer_email)
                                                  .state = Globals
                                                      .client.getCachedFriends[
                                                  index];
                                              selected_index = index;
                                            });
                                          }

                                          // remove friend.
                                          if (direction ==
                                              DismissDirection.startToEnd) {
                                            showDialog(
                                              context: context,
                                              builder: (_) => ConfirmDialog(
                                                danger: true,
                                                title:
                                                    'Delete ${Globals.client.getCachedFriends[index]}?',
                                                cancelFunction: () {
                                                  setState(() {
                                                    Globals.client
                                                        .getCachedFriends;
                                                  });
                                                  Navigator.of(context).pop();
                                                  return false;
                                                },
                                                actionText: Text('Approve'),
                                                actionFunction: () async {
                                                  try {
                                                    await Globals.client
                                                        .removeFriend(Globals
                                                                .client
                                                                .getCachedFriends[
                                                            index]);
                                                    setState(() {
                                                      Globals.client
                                                          .getCachedFriends
                                                          .removeAt(index);
                                                    });
                                                  } catch (e) {
                                                    ScaffoldMessenger.of(
                                                            context)
                                                        .showSnackBar(SnackBar(
                                                      content: Text(
                                                          '[${e.codeName}] ${e.message}'),
                                                      backgroundColor:
                                                          Colors.red,
                                                    ));
                                                  }

                                                  Navigator.of(context).pop();
                                                  return true;
                                                },
                                              ),
                                            );
                                          }

                                          return false;
                                        },
                                        secondaryBackground: Container(
                                          alignment: Alignment.centerRight,
                                          color: Colors.green,
                                          child: Icon(
                                              Icons.transfer_within_a_station),
                                        ),
                                        background: Container(
                                          alignment: Alignment.centerLeft,
                                          color: Colors.red,
                                          child: Icon(
                                              Icons.delete_forever_outlined),
                                        ),
                                        key: ValueKey(Globals
                                            .client.getCachedFriends[index]),
                                        child: ListTile(
                                          onTap: () async {
                                            setState(() {
                                              context
                                                  .read(Globals.transfer_email)
                                                  .state = Globals
                                                      .client.getCachedFriends[
                                                  index];
                                              selected_index = index;
                                            });
                                          },
                                          selected: selected_index == index,
                                          leading: CircleAvatar(
                                            backgroundImage: AssetImage(
                                                'assets/images/avatar.png'),
                                          ),
                                          dense: false,
                                          enabled: true,
                                          title: Text(
                                            "${Globals.client.getCachedFriends[index]}\n",
                                          ),
                                          subtitle: Text(
                                            "${Globals.client.getCachedFriends[index]}",
                                            style: TextStyle(
                                                color: Theme.of(context)
                                                    .highlightColor),
                                          ),
                                          trailing: Icon(
                                              Icons.transfer_within_a_station),
                                        ))),
                                separatorBuilder: (_, index) => SizedBox(
                                  height: 0,
                                ),
                                itemCount:
                                    Globals.client.getCachedFriends.length,
                              ),
                      ),
              ),
              index == 1
                  ? Padding(
                      padding: const EdgeInsets.only(left: 36.0, right: 36.0),
                      child: Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: <Widget>[
                            MyForm(
                              emailNode: _nodeText1,
                              amountNode: _nodeText2,
                            ),
                          ]),
                    )
                  : Container(),
            ],
          ),

          //Expanded(flex: 1, child: Container()),
        ],
      )),
    );
  }
}

class MyForm extends StatefulWidget {
  const MyForm({
    Key key,
    this.emailNode,
    this.amountNode,
  }) : super(key: key);

  final FocusNode emailNode;
  final FocusNode amountNode;

  @override
  _MyFormState createState() => _MyFormState();
}

class _MyFormState extends State<MyForm> {
  final emailControler = TextEditingController();
  final amountControler = TextEditingController();

  @override
  void dispose() {
    // Clean up the controller when the widget is disposed.
    emailControler.dispose();
    amountControler.dispose();
    super.dispose();
  }

  bool isNumeric(String s) {
    if (s == null) {
      return false;
    }
    return double.tryParse(s) != null;
  }

  final _formKey = GlobalKey<FormState>();
  bool _loading = false;

  @override
  Widget build(BuildContext context) {
    TextStyle style = TextStyle(
        fontFamily: 'Montserrat',
        fontSize: 20.0,
        color: Theme.of(context).highlightColor);

    emailControler.text = context.read(Globals.transfer_email).state;
    return Form(
      key: _formKey,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: <Widget>[
          TextFormField(
            focusNode: widget.emailNode,
            textInputAction: TextInputAction.next,
            keyboardType: TextInputType.emailAddress,
            controller: emailControler,
            style: style,
            decoration: InputDecoration(
                enabledBorder: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(32.0),
                  borderSide: BorderSide(color: Colors.grey, width: 1.0),
                ),
                contentPadding: EdgeInsets.fromLTRB(20.0, 15.0, 20.0, 15.0),
                hintText: "Email",
                hintStyle: TextStyle(color: Colors.grey),
                border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(32.0))),
            validator: (value) {
              if (value.isEmpty) {
                return 'Please enter the recipient email.';
              }
              return null;
            },
          ),
          SizedBox(height: 25.0),
          TextFormField(
            focusNode: widget.amountNode,
            textInputAction: TextInputAction.done,
            controller: amountControler,
            style: style,
            keyboardType: TextInputType.number,
            decoration: InputDecoration(
                enabledBorder: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(32.0),
                  borderSide: BorderSide(color: Colors.grey, width: 1.0),
                ),
                contentPadding: EdgeInsets.fromLTRB(20.0, 15.0, 20.0, 15.0),
                hintText: "Amount",
                hintStyle: TextStyle(color: Colors.grey),
                border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(32.0))),
            validator: (value) {
              if (value.isEmpty) {
                return 'Please enter the amount you would like to transfer';
              }

              if (!isNumeric(value)) {
                return 'Please enter a number';
              }

              if (num.parse(value) < 0) {
                return 'Amount must be above 0';
              }

              return null;
            },
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 16.0),
            child: Material(
              elevation: 5,
              borderRadius: BorderRadius.circular(30.0),
              color: Theme.of(context).highlightColor,
              child: MaterialButton(
                minWidth: MediaQuery.of(context).size.width,
                padding: EdgeInsets.fromLTRB(20.0, 15.0, 20.0, 15.0),
                onPressed: () async {
                  if (_formKey.currentState.validate()) {
                    showDialog(
                      context: context,
                      builder: (_) => ConfirmDialog(
                        title:
                            'Transfer ${amountControler.text} to ${emailControler.text}?',
                        cancelFunction: () {
                          Navigator.of(context).pop();
                          return false;
                        },
                        actionText: Text('Approve'),
                        actionFunction: () async {
                          try {
                            setState(() {
                              _loading = true;
                            });
                            await Globals.client.transferBalance(
                                emailControler.text,
                                int.parse(amountControler.text));

                            ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                              content: Text('Transfered successfully'),
                              backgroundColor: Colors.green,
                            ));

                            context.read(Globals.balance).state -=
                                int.parse(amountControler.text);
                            setState(() {
                              context.read(Globals.transfer_email).state = "";

                              _loading = false;
                            });
                          } on FormatException catch (e) {
                            setState(() {
                              _loading = false;
                            });
                            ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                              content:
                                  Text('${e.message}, is that number too big?'),
                              backgroundColor: Colors.red,
                            ));
                          } catch (e) {
                            setState(() {
                              _loading = false;
                            });
                            ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                              content: Text('[${e.codeName}] ${e.message}'),
                              backgroundColor: Colors.red,
                            ));
                          }

                          Navigator.of(context).pop();
                          return true;
                        },
                      ),
                    );
                  }
                },
                child: !_loading
                    ? Text("Transfer",
                        textAlign: TextAlign.center,
                        style: style.copyWith(
                            color: Theme.of(context).accentColor,
                            fontWeight: FontWeight.bold))
                    : //CircularProgressIndicator(valueColor: AlwaysStoppedAnimation<Color>(Colors.black),),
                    CupertinoActivityIndicator(),
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class SearchFriend extends StatefulWidget {
  @override
  _SearchFriendState createState() => _SearchFriendState();
}

class _SearchFriendState extends State<SearchFriend> {
  Future<List<SearchFriendResponse>> search(String search) async {
    await Future.delayed(Duration(seconds: 1));
    return await Globals.client.searchFriend(search.split('@')[0]).toList();
  }

  @override
  Widget build(BuildContext context) {
    return SearchBar<SearchFriendResponse>(
      searchBarStyle: SearchBarStyle(borderRadius: BorderRadius.circular(16.0)),
      minimumChars: 1,
      iconActiveColor: Theme.of(context).highlightColor,
      textStyle: TextStyle(color: Theme.of(context).highlightColor),
      cancellationText: Text(
        "clear",
        style: TextStyle(
            color: Theme.of(context).highlightColor,
            fontWeight: FontWeight.bold),
      ),
      onSearch: search,
      onItemFound: (SearchFriendResponse response, int index) {
        return ListTile(
          leading: CircleAvatar(
            backgroundImage: AssetImage('assets/images/avatar.png'),
          ),
          trailing: TextButton(
              onPressed: () async {
                showDialog(
                  context: context,
                  builder: (_) => ConfirmDialog(
                    title: 'Add ${response.mail}?',
                    cancelFunction: () {
                      Navigator.of(context).pop();
                      return false;
                    },
                    actionText: Text('Approve'),
                    actionFunction: () async {
                      try {
                        await Globals.client.addFriend(response.mail);
                        setState(() {
                          Globals.client.getCachedFriends.add(response.mail);
                        });

                        await Globals.client.getUserInfo();
                        ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                          content: Text('Added successfully'),
                          backgroundColor: Colors.green,
                        ));
                      } catch (e) {
                        ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                          content: Text('[${e.codeName}] ${e.message}'),
                          backgroundColor: Colors.red,
                        ));
                      }
                      Navigator.of(context).pop();
                      return true;
                    },
                  ),
                );
              },
              child: Text(
                "Add",
                style: TextStyle(color: Theme.of(context).highlightColor),
              )),
          title: Text(response.mail),
        );
      },
      onError: (error) {
        return Center(
          child: Text("Error occurred : $error"),
        );
      },
      loader: Center(
        // child: CircularProgressIndicator(
        //   valueColor: AlwaysStoppedAnimation<Color>(Theme.of(context).highlightColor),
        // ),
        child: CupertinoActivityIndicator(),
      ),
      emptyWidget: Center(
          child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          CircleAvatar(
              radius: 40,
              backgroundColor: Colors.grey.withOpacity(0.5),
              child: SizedBox(
                height: 50,
                child: Image.asset('assets/images/dog.png'),
              )),
          Text(
            "Wow, such empty",
            style: TextStyle(
                fontWeight: FontWeight.w600,
                color: Theme.of(context).highlightColor),
          )
        ],
      )),
    );
  }
}
