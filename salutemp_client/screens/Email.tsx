import React, { useState, useEffect } from 'react';
import { View, TextInput, Button, Alert, StyleSheet, Image, Dimensions, Text } from 'react-native';
import { FIREBASE_APP, FIREBASE_AUTH } from '../firebaseConfig';
import { signInWithEmailAndPassword, onAuthStateChanged, User } from 'firebase/auth';
import { useNavigation, StackActions } from '@react-navigation/native';
import colors from "../config/colors"
import { API_URL } from '../services/apiLinks';
import axios from "axios"

const Email = () => {

  const navigation = useNavigation();

  const [email, setEmail] = useState('');
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(FIREBASE_AUTH, (user) => {
      if (user) {
        setUser(user);
        console.log('signed in');
      } else {
        setUser(null);
        console.log('signed out');
      }
    });

    return unsubscribe;
  }, []);

  const getUserFromApi = async () => {
    await axios.get(`${API_URL}/v1/userexists/${email.toLocaleLowerCase()}`
    ).then((response) => {
      if (response.data["user"] == null)
      {
        console.log("no existing user");
        navigation.navigate("Name", {email: email.toLocaleLowerCase()});
      }
      else
      {
        console.log("existing user");
        navigation.navigate("EmailAndPassword", {email: email.toLocaleLowerCase()});
      }
    }).catch((error) => {
      console.log(`${API_URL}v1/userexists/${email}`);
      Alert.alert("Error Contacting Server", error.message)
    });
  };

  const handleContinue = async () => {
    if (email != "")
    {
      getUserFromApi()
      //navigation.navigate("EmailAndPassword", {email: email.toLocaleLowerCase()});
    }
    else
    {
      Alert.alert("Please Enter Email", "We need your email before continuing")
    }
  };

  if (user)
  {
    navigation.dispatch(StackActions.replace('MedList'));
  }
  else
  {
    return (
      <View style={styles.container}>
        <View style={styles.imagecontainer}>
          <Image
            source={require('../assets/salutemplogocolor.png')}
            style={styles.image}/>
        </View>
        <Text style={styles.text}>
          Log In or Sign Up
        </Text>
        <TextInput
          style={styles.input}
          placeholder="Email"
          value={email}
          onChangeText={(text) => setEmail(text)}/>
        <View style={styles.button}>
          <Button 
            title="Continue" 
            onPress={handleContinue} 
            color='#fff'/>
        </View>
      </View>
  );

  }

};

const styles = StyleSheet.create({
    container: {
      flex: 1,
      justifyContent: 'flex-start',
      alignItems: 'center',
      backgroundColor: colors.background,
    },
    imagecontainer: {
        justifyContent: 'flex-start',
        alignItems: 'center',
        backgroundColor: colors.background,
        width: '30%',
        height: '20%',
        margin: 10,
        marginTop: 20
    },
    input: {
      width: '80%',
      height: 40,
      paddingHorizontal: 10,
      margin: 10,
      borderWidth: 1,
      borderRadius: 10,
      backgroundColor: colors.grey,
      borderColor: colors.grey,
    },
    image: {
        flex: 1,
        width: '100%',
        height: '100%',
        borderBottomLeftRadius: 20,
        borderBottomRightRadius: 20,
    },
    button: {
        backgroundColor:colors.darkNeutral,
        borderRadius: 10,
        width: '80%',
        margin: 10,
    },
    text: {
      fontSize: 20,
      color: colors.darkNeutral,
      fontWeight: "700",
      margin: 10,
      marginBottom: 75
    }
  });

export default Email;
