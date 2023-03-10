import "./Login.css"
import React, {useState} from "react";
import putModel from "../middleware/putData";
import {DASHBOARD_URL, MANAGER_URL} from "../constants/WebURL";


/**
 * Login screen
 * @returns {JSX.Element}
 * @constructor
 */
function Login() {
    const [credentials, setCredentials] = useState({
        userName : "",
        password : ""
    })

    /**
     * Function that will check the credentials.
     *
     * @returns {Promise<void>}
     */
    const login = async () => {
        await putModel(MANAGER_URL, credentials)
            .then(res=> {
                sessionStorage.setItem("credentials", (res.text))
               window.location.href = DASHBOARD_URL
            })
            .catch(res => console.log(res))
    }




    return (
        <div className="login-page" style={{display: "flex", flexDirection: "column"}}>
            <input type={"text"}
                   placeholder={"Username"}
                   name={"userName"}
                   value={credentials.userName}
                   onChange={(e) => {
                       setCredentials({...credentials, userName: e.target.value})
                   }}/>


            <input type={"password"}
                   placeholder={"Password"}
                   name={"password"}
                   value={credentials.password}
                   onChange={(e) => {
                       setCredentials({...credentials, password: e.target.value})
                   }}/>

            <button className={"btn"} onClick={login}>Login</button>
        </div>
    );
}

export default Login;
