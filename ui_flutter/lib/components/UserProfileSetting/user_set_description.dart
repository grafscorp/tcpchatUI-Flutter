import 'package:flutter/material.dart';

class UserSetDescription extends StatefulWidget {
  const UserSetDescription({super.key});

  @override
  State<UserSetDescription> createState() => _UserSetDescriptionState();
}

class _UserSetDescriptionState extends State<UserSetDescription> {
  final _textDescriptionController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: _textDescriptionController,
      maxLines: 10,
      decoration: InputDecoration(
        hintText: "Enter description...",
        labelText: "Description",
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(40)),
      ),
      onChanged: (val) {},
      onSubmitted: (val) {},
    );
  }
}
