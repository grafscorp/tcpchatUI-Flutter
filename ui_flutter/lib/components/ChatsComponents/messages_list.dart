import 'package:flutter/material.dart';
import 'package:ui_flutter/components/ChatsComponents/message_box.dart';
import 'package:ui_flutter/components/ChatsComponents/server_message_box.dart';

class MessagesLists extends StatefulWidget {
  const MessagesLists({super.key});

  @override
  State<MessagesLists> createState() => _MessagesListsState();
}

class _MessagesListsState extends State<MessagesLists> {
  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.only(bottom: 20),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.end,
        crossAxisAlignment: CrossAxisAlignment.start,
        mainAxisSize: MainAxisSize.max,
        children: [
          ServerMessageBox(text: "Meida is connected"),
          MessageBox(
            nickname: "Meida",
            messageText: "Hello",
            timesend: "00:40",
            myMessage: false,
          ),
          MessageBox(
            nickname: "Hecker",
            messageText: "Bye bye",
            timesend: "00:41",
            myMessage: true,
          ),
        ],
      ),
    );
  }
}
