import 'package:flutter/material.dart';

class UserSetAvatar extends StatefulWidget {
  const UserSetAvatar({super.key});

  @override
  State<UserSetAvatar> createState() => _UserSetAvatarState();
}

class _UserSetAvatarState extends State<UserSetAvatar> {
  @override
  Widget build(BuildContext context) {
    return Container(
      width: 150,
      height: 150,
      decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(75), color: Colors.amber),
    );
  }
}
