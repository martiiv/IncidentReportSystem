import {useEffect, useState} from "react";
import {MANAGER_URL} from "../constants/WebURL";
import fetchData from "../middleware/FetchData";

function Profile() {

    const [profileData, setProfileData] = useState("")
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        const fetch = async () => {
            const data = await fetchData(MANAGER_URL + "?id=1");
            setProfileData(data.data)

        }

        fetch().then(() => setLoading(false))
            // make sure to catch any error
            .catch(console.error);
    }, [])




    console.log(profileData)
    if (loading){
        return (
            <h1>Loading</h1>
        )
    }else {
        return (
            <div className="App">
            <span>
                <p>Name: {profileData.userName}</p>
                <p>Company: {profileData.company}</p>

            </span>
            </div>

        );
    }

}

export default Profile;
