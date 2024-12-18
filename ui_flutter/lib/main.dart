import 'package:flutter/material.dart';
import 'package:ui_flutter/pages/group_chat_page.dart';
import 'package:ui_flutter/pages/home_page.dart';
import 'package:ui_flutter/pages/test_chat_page.dart';

void main() {
  runApp(const MainApp());
}

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "TCPChat",
      theme: ThemeData.dark(),
      home: Scaffold(
        body: Center(
          child: TestChatPage(),
        ),
      ),
    );
  }
}
