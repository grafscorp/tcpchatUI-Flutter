import 'package:flutter/material.dart';

class GroupChatPage extends StatefulWidget {
  const GroupChatPage({super.key});

  @override
  State<GroupChatPage> createState() => _GroupChatPageState();
}

class _GroupChatPageState extends State<GroupChatPage> {
  double screenWidth = 0;
  final primColor = const Color.fromARGB(255, 44, 39, 52);
  final secColor = const Color.fromARGB(255, 25, 21, 32);
  @override
  Widget build(BuildContext context) {
    screenWidth = MediaQuery.sizeOf(context).width;
    return Scaffold(
      drawer: Drawer(
        child: Text("Test"),
      ),
      body: Row(
        children: [
          Container(
            width: screenWidth > 350 ? 350 : 70,
            color: Color.fromARGB(255, 44, 39, 52),
            child: Column(
              children: [
                Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: Row(
                    children: [
                      IconButton(
                        onPressed: () {
                          Scaffold.of(context).openDrawer();
                        },
                        icon: Icon(
                          Icons.menu,
                          size: 40,
                        ),
                      )
                    ],
                  ),
                ),
              ],
            ),
          ),
          Text("data"),
        ],
      ),
    );
  }
}
