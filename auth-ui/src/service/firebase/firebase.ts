// import firebase from 'firebase/app';
import 'firebase/database';
import * as firebase from "firebase/app";
import { getAuth, GoogleAuthProvider  } from "firebase/auth";

const firebaseConfig = {
    apiKey: process.env.NEXT_PUBLIC_REACT_APP_FIREBASE_API_KEY?.trim(),
    authDomain: process.env.NEXT_PUBLIC_REACT_APP_FIREBASE_AUTH_DOMAIN?.trim(),
    projectId: process.env.NEXT_PUBLIC_REACT_APP_FIREBASE_PROJECT_ID?.trim(),
};

const firebaseApp = firebase.initializeApp(firebaseConfig);

export const firebaseAuth = getAuth(firebaseApp);
export default firebaseApp;

export const googleProvider = new GoogleAuthProvider();