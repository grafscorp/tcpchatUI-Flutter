import 'package:flutter/material.dart';

class UserSetName extends StatefulWidget {
  const UserSetName({super.key});

  @override
  State<UserSetName> createState() => _UserSetNameState();
}

class _UserSetNameState extends State<UserSetName> {
  @override
  Widget build(BuildContext context) {
    return TextField(
      decoration: InputDecoration(
        icon: Icon(Icons.person),
        labelText: "Nickname",
        border: UnderlineInputBorder(),
      ),
      onChanged: (val) {},
      onSubmitted: (val) {},
    );
  }
}
