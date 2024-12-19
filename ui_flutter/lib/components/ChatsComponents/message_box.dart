import 'package:flutter/material.dart';

class MessageBox extends StatelessWidget {
  MessageBox(
      {super.key,
      required this.messageText,
      required this.nickname,
      required this.timesend,
      required this.myMessage});
  String messageText;
  String nickname;
  String timesend;
  bool myMessage = true;

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment:
          !myMessage ? MainAxisAlignment.start : MainAxisAlignment.end,
      crossAxisAlignment: CrossAxisAlignment.end,
      children: [
        //Avatar
        !myMessage
            ? Column(
                children: [
                  Container(
                    padding: EdgeInsets.only(bottom: 20),
                    decoration: BoxDecoration(
                        color: Colors.grey,
                        borderRadius: BorderRadius.circular(20)),
                    width: 40,
                    height: 40,
                  ),
                  SizedBox(
                    height: 5,
                  ),
                ],
              )
            : Container(),
        SizedBox(
          width: 6,
        ),
        Column(
          crossAxisAlignment:
              !myMessage ? CrossAxisAlignment.start : CrossAxisAlignment.end,
          children: [
            //Nickname
            Text(nickname),
            Container(
              height: 80,
              width: 180,
              decoration: BoxDecoration(
                color: !myMessage ? Colors.lightBlueAccent : Colors.white70,
                borderRadius: BorderRadius.only(
                  topLeft: Radius.circular(20),
                  topRight: Radius.circular(20),
                  bottomLeft:
                      !myMessage ? Radius.circular(0) : Radius.circular(20),
                  bottomRight:
                      !myMessage ? Radius.circular(20) : Radius.circular(0),
                ),
              ),
              child: Padding(
                padding: const EdgeInsets.all(8.0),
                child: Text(messageText),
              ),
            ),
            //Date time
            Text(timesend),
          ],
        ),
      ],
    );
  }
}
