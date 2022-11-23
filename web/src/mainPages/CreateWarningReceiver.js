import React, {useEffect, useState} from "react";
import postModel from "../middleware/postData";
import {GROUPS_URL, RECEIVER_URL} from "../constants/WebURL";
import Select from 'react-select'
import "./Create.css"
import fetchData from "../middleware/FetchData";
import {useNavigate} from "react-router-dom";
import {ReactNotifications, Store} from "react-notifications-component";

/**
 * Page to add new warning receiver.
 * @returns {JSX.Element}
 * @constructor
 */
function CreateWarningReceiver(){
    const [warning, setWarning] = useState({
        name: "",
        phoneNumber: "",
        company: "",
        receiverGroup: "",
        receiverEmail: ""
    })

    const [options, setOptions] = useState([])
    const navigate = useNavigate()

    /**
     * Function that will submit the warning receiver to the database.
     * Will Send confirmation notification on success and rejection.
     *
     * @param event data sent from forms
     * @returns {Promise<void>}
     */
    const handleSubmit = async (event) => {
        event.preventDefault();
        await postModel(RECEIVER_URL, JSON.stringify(warning))
            .then(() => {
                Store.addNotification({
                    title: "Successfully Created!",
                    message: "New Warning Receiver created",
                    type: "success",
                    insert: "top",
                    container: "top-right",
                    animationIn: ["animate__animated", "animate__fadeIn"],
                    animationOut: ["animate__animated", "animate__fadeOut"],
                    dismiss: {
                        duration: 2000,
                        onScreen: true
                    }
                });
            })
            .then(setTimeout(() => {
            navigate(-1)
        }, 2000))
            .catch(() => alert("something went wrong. please try again"))
    }

    /**
     * Function that will be called when the page is loaded.
     * Function will fetch data of group.
     */
    useEffect(() => {
        const fetch = async () => {
            const data = await fetchData(GROUPS_URL)
            setOptions(data.data.map(item => ({value: item.name, label: item.name})))
        }
        fetch().catch(() => alert("No connection, try again later"))
    }, [])


    /**
     * Function that will set the selected receiver group to warning receiver.
     * @param e data from input field.
     */
    function handleChangeOptions(e) {
        setWarning({
            ...warning, "receiverGroup": e.value
        })
    }


    /**
     * Function that will set the given value to warning receiver.
     * @param evt from input field.
     * */
    function handleChange(evt) {
        const value = evt.target.value;
        setWarning({
            ...warning, [evt.target.name]: value
        })
    }


    return(
        <div className={"create"}>
            <h2>New Warning Receiver</h2>
            <ReactNotifications />
            <form className={"create-forms"} onSubmit={handleSubmit} onKeyDown={(e) => { e.key === 'Enter' && e.preventDefault(); }}>
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
                    <Select options={options} onChange={handleChangeOptions} className={"input-group select-group"}/>
                </label>

                <label>Email:
                    <input
                        type={"email"}
                        name={"receiverEmail"}
                        value={warning.receiverEmail}
                        onChange={handleChange}
                    />
                </label>
                <div className={"btn-group"}>
                <button className={"btn"} onClick={() => navigate(-1)}>Back</button>
                <button className={"btn send-btn"}>SEND</button>
                </div>
            </form>
        </div>
    )
}
export default CreateWarningReceiver
