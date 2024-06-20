import { View, Text } from "react-native";
import React, { useEffect, useState } from "react";
import { getFirestore, getDocs, collection } from "firebase/firestore";

export default function AddPostScreen() {
  const db = getFirestore(app);

  const [categoryList, setCategoryList] = useState([]);
  useEffect(() => {
    getCategooryList;
  }, []);

  /*Firebase code to retrive all documents in Category collection*/
  const getCategooryList = async () => {
    const querySnapshot = await getDocs(collection(db, "Category"));
    querySnapshot.forEach((doc) => {
      //console.log(doc.data());

      /*Append the categoryList with */
      setCategoryList((categoryList) => [...categoryList, doc.data()]);
    });
  };
  return (
    <View>
      <Text>AddPostScreen</Text>
    </View>
  );
}
