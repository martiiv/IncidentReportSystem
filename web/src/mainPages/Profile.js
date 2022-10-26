import {useEffect, useState} from "react";
import {MANAGER_URL} from "../constants/WebURL";
import fetchData from "../middleware/FetchData";

function Profile() {

    const [profileData, setProfileData] = useState("")

    useEffect(() => {
        const fetch = async () => {
            const data = await fetchData(MANAGER_URL + "?id=1");
            setProfileData(data)
        }

        fetch()
            // make sure to catch any error
            .catch(console.error);
    }, [])




    console.log(profileData.data)
    return (
        <div className="App">
            <span>
                <p>Name: {profileData.data.userName}</p>
                <p>Company: {profileData.data.company}</p>

            </span>
        </div>
    );
}

export default Profile;
