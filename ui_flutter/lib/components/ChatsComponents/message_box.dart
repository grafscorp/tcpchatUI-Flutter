import 'package:flutter/material.dart';

class MessageBox extends StatelessWidget {
  MessageBox({super.key, required this.messageText});
  String messageText;

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 80,
      width: 180,
      decoration: BoxDecoration(
        color: Colors.lightBlueAccent,
        borderRadius: BorderRadius.only(
          topLeft: Radius.circular(20),
          topRight: Radius.circular(20),
          bottomLeft: Radius.circular(0),
          bottomRight: Radius.circular(20),
        ),
      ),
    );
  }
}
