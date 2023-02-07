import {firebaseAuth,  googleProvider} from './firebase'; 
import {
    signInWithPopup,
  } from "firebase/auth";

class Auth {
    async login(name:any) {
        try {
            const provider = this.getProvider(name);
            return await signInWithPopup(firebaseAuth ,provider).then()
            
        } catch (error) {
            console.error(error)
            return null
        }
    }
    async logout() {
        try {
            await firebaseAuth.signOut()

        } catch (error) {
            alert("error! ," + error)
        }
    }
    getProvider(name:any){
        switch(name){
          case 'Google':
            return googleProvider;
          default:
            throw new Error(`${name} is unknown provider.`);
        }
    }
    onAuthChange = (callback:any) => {
        firebaseAuth.onAuthStateChanged(user => {
            callback(user)
        }) 
    }
};
export default Auth;