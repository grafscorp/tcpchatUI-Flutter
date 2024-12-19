import 'package:flutter/material.dart';
import 'package:ui_flutter/components/ChatsComponents/message_box.dart';

class ServerMessageBox extends StatelessWidget {
  ServerMessageBox({super.key, required this.text});
  String text;
  @override
  Widget build(BuildContext context) {
    return Center(
      child: Container(
        decoration: BoxDecoration(
          color: const Color.fromARGB(179, 46, 43, 43),
          borderRadius: BorderRadius.circular(20),
          border: Border.all(
            width: 2,
            color: Color.fromARGB(255, 255, 255, 255),
          ),
        ),
        height: 30,
        width: 400,
        child: Center(child: Text(text)),
      ),
    );
  }
}
