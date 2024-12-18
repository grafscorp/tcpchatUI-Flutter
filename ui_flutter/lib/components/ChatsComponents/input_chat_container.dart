import 'package:flutter/material.dart';

class InputChatContainer extends StatefulWidget {
  const InputChatContainer({super.key});

  @override
  State<InputChatContainer> createState() => _InputChatContainerState();
}

class _InputChatContainerState extends State<InputChatContainer> {
  final inputIconSize = 30.0;
  @override
  Widget build(BuildContext context) {
    return Container(
      //Input Container
      decoration: BoxDecoration(
        color: Colors.grey,
        borderRadius: BorderRadius.circular(20),
      ),
      height: 70,
      child: Padding(
        padding: const EdgeInsets.all(4.0),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Row(
              children: [
                //Upload file button
                IconButton(
                  onPressed: () {},
                  icon: Icon(
                    Icons.upload_file_outlined,
                    size: inputIconSize,
                  ),
                ),
                //Text Input
              ],
            ),
            IconButton(
              onPressed: () {},
              icon: Icon(
                Icons.send_rounded,
                size: inputIconSize,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
