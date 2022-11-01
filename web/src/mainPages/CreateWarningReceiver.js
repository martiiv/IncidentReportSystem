import {useState} from "react";
import postModel from "../middleware/postData";
import {RECEIVER_URL} from "../constants/WebURL";

import "./CreateWarningReceiver.css"


function CreateWarningReceiver(){

    const [warning, setWarning] = useState({
        name: "",
        phoneNumber: "",
        company: "",
        receiverGroup: "",
        receiverEmail: ""
    })

    const handleSubmit = async (event) => {
        event.preventDefault();
        await postModel(RECEIVER_URL, warning)
            .then(res => console.log(res))

    }


    function handleChange(evt) {
        const value = evt.target.value;
        setWarning({
            ...warning, [evt.target.name]: value
        })
    }

    console.log(warning)
    return(
        <div className={"create-warning-receiver"}>
            <h2>New Warning Receiver</h2>
            <form className={"create-warning-receiver"} onSubmit={handleSubmit} onKeyDown={(e) => { e.key === 'Enter' && e.preventDefault(); }}>
                <label>Name:
                    <input
                        type={"text"}
                        name={"name"}
                        required
                        value={warning.name}
                        onChange={handleChange}
                    />
                </label>

                <label>Phone:
                    <input
                        type={"text"}
                        name={"phoneNumber"}
                        value={warning.phoneNumber}
                        onChange={handleChange}
                    />
                </label>

                <label>Comapny:
                    <input
                        type={"text"}
                        name={"company"}
                        required
                        value={warning.company}
                        onChange={handleChange}
                    />
                </label>

                <label>Receiver Group:
                    <textarea
                        name={"receiverGroup"}
                        required
                        value={warning.receiverGroup}
                        onChange={handleChange}
                    />
                </label>

                <label>Email:
                    <input
                        type={"text"}
                        name={"receiverEmail"}
                        value={warning.receiverEmail}
                        onChange={handleChange}
                    />
                </label>
             <button className={"btn"}>SEND </button>
            </form>
        </div>
    )
}
export default CreateWarningReceiver
