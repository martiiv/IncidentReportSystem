import {useEffect, useState} from "react";
import postModel from "../middleware/postData";
import {GROUPS_URL, RECEIVER_URL} from "../constants/WebURL";
import Select from 'react-select'
import "./CreateWarningReceiver.css"
import fetchData from "../middleware/FetchData";


function CreateWarningReceiver(){

    const [warning, setWarning] = useState({
        name: "",
        phoneNumber: "",
        company: "",
        receiverGroup: "",
        receiverEmail: ""
    })

    const [options, setOptions] = useState([])

    const handleSubmit = async (event) => {
        event.preventDefault();
        await postModel(RECEIVER_URL, warning)
            .then(res => console.log(res))

    }

    useEffect(() => {
        const fetch = async () => {
            const data = await fetchData(GROUPS_URL)
            console.log(data.data)

            setOptions(data.data.map(item => ({value: item.name, label: item.name})))
        }

        fetch().catch(e => console.log(e))
    }, [])



    function handleChangeOptions(e) {
       const group = (e.map(item => item.value).join(","))
        setWarning({
            ...warning, ["receiverGroup"]: group
        })
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

                <label>Company:
                    <input
                        type={"text"}
                        name={"company"}
                        required
                        value={warning.company}
                        onChange={handleChange}
                    />
                </label>

                <label>Receiver Group:
                    <Select isMulti options={options} onChange={handleChangeOptions} className={"input-group"}/>
                </label>

                <label>Email:
                    <input
                        type={"email"}
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
