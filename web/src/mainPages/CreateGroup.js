import {useState} from "react";
import postModel from "../middleware/postData";
import {GROUPS_URL, RECEIVER_URL} from "../constants/WebURL";
import {useNavigate} from "react-router-dom";


function CreateGroup(){

    const [group, setGroup] = useState({
        name: "",
        info: "",
    })

    const [success, setSuccess] = useState(true)

    const navigate = useNavigate()

    const handleSubmit = async (event) => {
        event.preventDefault();
        await postModel(GROUPS_URL, group)
            .then(() => navigate(-1))
            .catch(() => setSuccess(false))
    }

    function handleChange(evt) {
        const value = evt.target.value;
        setGroup({
            ...group, [evt.target.name]: value
        })
    }

    console.log(group)


    return(
        <div className={"create-warning-receiver"}>
            {success ? null : <h1>Not added</h1>}
            <h2>New Group</h2>
            <form className={"create-warning-receiver"} onSubmit={handleSubmit} onKeyDown={(e) => { e.key === 'Enter' && e.preventDefault(); }}>
                <label>Name:
                    <input
                        type={"text"}
                        name={"name"}
                        required
                        value={group.name}
                        onChange={handleChange}
                    />
                </label>

                <label>Information:
                    <input
                        type={"text"}
                        name={"info"}
                        value={group.info}
                        onChange={handleChange}
                    />
                </label>
                <button className={"btn"}>SEND</button>
            </form>
        </div>
    )
}
export default CreateGroup
