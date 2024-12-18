import 'package:flutter/material.dart';
import 'package:ui_flutter/components/ChatsComponents/input_chat_container.dart';
import 'package:ui_flutter/components/ChatsComponents/message_box.dart';

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
          mainAxisAlignment: MainAxisAlignment.end,
          children: [
            MessageBox(
              messageText: "Test",
            ),
            InputChatContainer(),
          ],
        ),
      ),
    );
  }
}
