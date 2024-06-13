import { StatusBar } from "expo-status-bar";
import { StyleSheet, Text, View } from "react-native";
import LoginScreen from "./Apps/Screens/LoginScreen";
import { ClerkProvider, SignedIn, SignedOut } from "@clerk/clerk-expo";

export default function App() {
  return (
    <ClerkProvider publishableKey="pk_test_c3dlZXBpbmctZ2xpZGVyLTQxLmNsZXJrLmFjY291bnRzLmRldiQ">
      <View className="flex-1 items-center justify-center bg-white">
        <SignedIn>
          <Text>You are Signed in</Text>
        </SignedIn>
        <SignedOut>
          <LoginScreen />
        </SignedOut>
      </View>
    </ClerkProvider>
  );
}
