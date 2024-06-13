import React from "react";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import HomeScreen from "../Screens/HomeScreen";
import ProfileScreen from "../Screens/ProfileScreen";
import AddPostScreen from "../Screens/AddPostScreen";
import ExploreScreen from "../Screens/ExploreScreen";
import { AntDesign } from "@expo/vector-icons";
import { Text } from "react-native";

const Tab = createBottomTabNavigator();

export default function TabNavigation() {
  return (
    <Tab.Navigator
      screenOptions={{ headerShown: false, tabBarActiveTintColor: "#000" }}
    >
      <Tab.Screen
        name="Home"
        options={{
          tabBarLabel: ({ color }) => (
            <Text style={{ color: color }}> Home</Text>
          ),
          tabBarIcon: ({ color, size }) => (
            <AntDesign name="home" size={24} color={color} />
          ),
        }}
        component={HomeScreen}
      />
      <Tab.Screen
        name="Explore"
        options={{
          tabBarLabel: ({ color }) => (
            <Text style={{ color: color }}> Explore</Text>
          ),
          tabBarIcon: ({ color, size }) => (
            <AntDesign name="search1" size={24} color={color} />
          ),
        }}
        component={ExploreScreen}
      />
      <Tab.Screen
        name="AddPost"
        options={{
          tabBarLabel: ({ color }) => (
            <Text style={{ color: color }}> Add </Text>
          ),
          tabBarIcon: ({ color, size }) => (
            <AntDesign name="plussquareo" size={24} color={color} />
          ),
        }}
        component={AddPostScreen}
      />
      <Tab.Screen
        name="Profile"
        options={{
          tabBarLabel: ({ color }) => (
            <Text style={{ color: color }}> Profile</Text>
          ),
          tabBarIcon: ({ color, size }) => (
            <AntDesign name="user" size={24} color="black" />
          ),
        }}
        component={ProfileScreen}
      />
    </Tab.Navigator>
  );
}
