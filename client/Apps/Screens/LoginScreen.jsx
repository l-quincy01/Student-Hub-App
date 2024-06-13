import {
  View,
  Text,
  Image,
  Dimensions,
  SafeAreaView,
  TouchableOpacity,
} from "react-native";
import React from "react";

import * as WebBrowser from "expo-web-browser";
import { useWarmUpBrowser } from "../../hooks/warmUpBrowser";
import { useOAuth } from "@clerk/clerk-expo";
WebBrowser.maybeCompleteAuthSession();

export default function LoginScreen() {
  useWarmUpBrowser();
  const { width, height } = Dimensions.get("window");
  const { startOAuthFlow } = useOAuth({ strategy: "oauth_google" });

  const onPress = React.useCallback(async () => {
    try {
      const { createdSessionId, signIn, signUp, setActive } =
        await startOAuthFlow();

      if (createdSessionId) {
        setActive({ session: createdSessionId });
      } else {
        // Use signIn or signUp for next steps such as MFA
      }
    } catch (err) {
      console.error("OAuth error", err);
    }
  }, []);

  return (
    <SafeAreaView style={{ flex: 1 }}>
      {/* <Text className="text-red-500">LoginScreen</Text> */}
      <Image
        source={require("./../../assets/images/banner.jpeg")}
        style={{
          top: 20,
          width: width,
          height: height / 4,
          resizeMode: "cover",
        }}
      />
      <View className="p-10 bg-white mt-[-5px] rounded-t-3xl shadow-lg">
        <Text className="text-[30px] font-bold"> College Marketplace</Text>
        <Text className="text-[18px] text-slate-500 mt-5">
          {" "}
          Buy and sell from other students in your college
        </Text>
        <TouchableOpacity
          onPress={onPress}
          className=" mt-20 p-4 bg-blue-400 rounded-xl "
        >
          <Text className="text-white text-center text-[18px]">
            Get Started
          </Text>
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
}
