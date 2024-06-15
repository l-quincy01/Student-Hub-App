import {
  View,
  Text,
  Image,
  Dimensions,
  SafeAreaView,
  TouchableOpacity,
} from "react-native";
import React from "react";
import { AntDesign } from "@expo/vector-icons";
import { Icon } from "@rneui/themed";

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
    <View className="flex-1  bg-white ">
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
      <View className=" fon gap-y-1 p-5 rounded-t-3xl shadow-lg gro">
        <View className="p-3 gap-y-1">
          <Text className=" text-2xl text-center font-bold">
            College Marketplace
          </Text>
          <Text className=" border-c text-md text-slate-500  text-center font-semibold">
            Buy and sell from other students in your college.
          </Text>
        </View>

        {/* Login button options */}
        <View className=" flex-col gap-y-4 px-5 mb-3 ">
          <TouchableOpacity
            onPress={onPress}
            className=" p-4 border-2 border-gray-300 rounded-md items-center justify-center flex-row gap-x-2"
          >
            <Text className="font-semibold text-black text-center text-md">
              Sign Up With Email
            </Text>
          </TouchableOpacity>
          <TouchableOpacity
            onPress={onPress}
            className=" p-4 border-2 border-gray-300 rounded-md items-center justify-center flex-row gap-x-2"
          >
            <Text className="font-semibold text-black text-center text-md">
              Continue With Email
            </Text>
          </TouchableOpacity>
        </View>

        <Text className=" border border-black text-center"> or</Text>

        {/*Third party sign-up options */}
        <View className="flex-col gap-y-4 px-5">
          <TouchableOpacity
            onPress={onPress}
            className=" p-4 border-2 border-gray-300 rounded-md items-center justify-center flex-row gap-x-2"
          >
            <AntDesign name="google" size={16} color="black" />
            <Text className="font-semibold text-black text-center text-md">
              Continue with Google
            </Text>
          </TouchableOpacity>
          <TouchableOpacity
            onPress={onPress}
            className="bg-black p-4 border-2 border-black  rounded-md items-center justify-center flex-row gap-x-2"
          >
            <AntDesign name="apple1" size={16} color="white" />
            <Text className="font-semibold text-white text-center text-md">
              Continue with Apple
            </Text>
          </TouchableOpacity>
        </View>
      </View>
    </View>
  );
}
