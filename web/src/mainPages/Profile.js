import {useEffect, useState} from "react";
import {MANAGER_URL} from "../constants/WebURL";
import fetchData from "../middleware/FetchData";

function Profile() {

        const credentials = JSON.parse(sessionStorage.getItem("credentials"))
        return (
            <div className="App">
            <span>
                <p>Name: {credentials.userName}</p>
                <p>Company: {credentials.company}</p>

            </span>
            </div>

        );


}

export default Profile;
