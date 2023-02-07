import React from "react";
import { GoogleLoginButton } from "react-social-login-buttons";

const Login = ({auth}:any) => {
    const onLogin = (e:any) => {
        auth.login(e.target.textContent).then(console.log);
    }

    return (
        <div>
            <li>
                {/* <button onClick={onLogin}>Google</button> */}
                <GoogleLoginButton onClick={() => {
                    onLogin("Google");
                }}></GoogleLoginButton>
            </li>
        </div>
    )
}

export default Login