import "./Login.css"

function Login() {
    return (
        <div className="login-page" style={{display: "flex", flexDirection: "column"}}>
            <input type={"text"} placeholder={"Username"}/>
            <input type={"password"} placeholder={"Password"}/>
            <button className={"btn"}>Login</button>
        </div>
    );
}

export default Login;
