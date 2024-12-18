import 'package:flutter/material.dart';

class GroupChatPage extends StatefulWidget {
  const GroupChatPage({super.key});

  @override
  State<GroupChatPage> createState() => _GroupChatPageState();
}

class _GroupChatPageState extends State<GroupChatPage> {
  double screenWidth = 0;

  @override
  Widget build(BuildContext context) {
    screenWidth = MediaQuery.sizeOf(context).width;
    return Scaffold(
      body: Row(
        children: [
          Container(
            width: screenWidth > 350 ? 350 : 70,
            color: Color.fromARGB(255, 44, 39, 52),
          ),
          Expanded(
            child: Container(
              color: Color.fromARGB(255, 25, 21, 32),
              width: 600,
            ),
          )
        ],
      ),
    );
  }
}
