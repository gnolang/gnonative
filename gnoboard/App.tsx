import React, { useEffect } from "react";
import { StyleSheet } from "react-native";

import { GoBridge } from "@gno/native_modules/GoBridge";
import HomeScreen from "@gno/screens/home";
import SettingsScreen from "@gno/screens/settings";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { NavigationContainer } from "@react-navigation/native";

const Tab = createBottomTabNavigator();

export default function App() {
  const [greeting, setGreeting] = React.useState<string>("");

  useEffect(() => {
    const getGreeting = async () =>
      setGreeting(await GoBridge.getAccountAndTxCfg());

    getGreeting();
  }, []);

  return (
    <NavigationContainer>
      <Tab.Navigator>
        <Tab.Screen name="Home" component={HomeScreen} />
        <Tab.Screen name="Settings" component={SettingsScreen} />
      </Tab.Navigator>
    </NavigationContainer>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
});
