import 'package:flutter/material.dart';
import 'package:ui_flutter/components/ChatsComponents/input_chat_container.dart';
import 'package:ui_flutter/components/ChatsComponents/message_box.dart';
import 'package:ui_flutter/components/ChatsComponents/messages_list.dart';

class TestChatPage extends StatefulWidget {
  const TestChatPage({super.key});

  @override
  State<TestChatPage> createState() => _TestChatPageState();
}

class _TestChatPageState extends State<TestChatPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("TCP Chat"),
        centerTitle: true,
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          mainAxisSize: MainAxisSize.max,
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisAlignment: MainAxisAlignment.end,
          children: [
            Expanded(child: MessagesLists()),
            InputChatContainer(),
          ],
        ),
      ),
    );
  }
}
