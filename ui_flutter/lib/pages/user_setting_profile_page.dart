import 'package:flutter/material.dart';
import 'package:ui_flutter/components/UserProfileSetting/user_set_avatar.dart';
import 'package:ui_flutter/components/UserProfileSetting/user_set_description.dart';
import 'package:ui_flutter/components/UserProfileSetting/user_set_name.dart';

class UserSettingProfile extends StatefulWidget {
  const UserSettingProfile({super.key});

  @override
  State<UserSettingProfile> createState() => _UserSettingProfileState();
}

class _UserSettingProfileState extends State<UserSettingProfile> {
  Size btnSize = Size(140, 50);

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("My Profile"),
      ),
      body: Padding(
        padding: EdgeInsets.symmetric(vertical: 40, horizontal: 60),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            //*Test
            //User Avatar
            UserSetAvatar(),
            SizedBox(
              height: 20,
            ),
            //UserName
            SizedBox(width: 500, child: UserSetName()),
            SizedBox(
              height: 40,
            ),
            //User Description
            SizedBox(width: 500, child: UserSetDescription()),
            SizedBox(
              height: 60,
            ),
            //Buttons
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                SizedBox(
                  height: btnSize.height,
                  width: btnSize.width,
                  child: ElevatedButton(child: Text("Back"), onPressed: () {}),
                ),
                SizedBox(
                  width: 90,
                ),
                SizedBox(
                  height: btnSize.height,
                  width: btnSize.width,
                  child: ElevatedButton(child: Text("Save"), onPressed: () {}),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
