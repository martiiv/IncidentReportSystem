import React, {useState} from "react";
import postModel from "../middleware/postData";
import {GROUPS_URL} from "../constants/WebURL";
import {useNavigate} from "react-router-dom";
import "./Create.css"
import {Store} from "react-notifications-component";


/**
 * Page that will make it possible for the user to create a new Group
 * @returns {JSX.Element}
 * @constructor
 */
function CreateGroup() {
    const [group, setGroup] = useState({
        name: "",
        info: "",
    })

    const [success, setSuccess] = useState(true)
    const navigate = useNavigate()

    /**
     * Function that will submit the new group into the database.
     * Will Send confirmation notification on success and rejection.
     *
     * @param event data to be submitted.
     * @returns {Promise<void>}
     */
    const handleSubmit = async (event) => {
        event.preventDefault();
        await postModel(GROUPS_URL, JSON.stringify(group))
            .then(() => {
                Store.addNotification({
                    title: "Successfully Created!",
                    message: "New Group created",
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
            .catch(() => {
                setSuccess(false)
                alert("something went wrong. please try again")
            })
    }

    /**
     * Will set the group name and information.
     * @param evt data sent from input field.
     */
    function handleChange(evt) {
        const value = evt.target.value;
        setGroup({
            ...group, [evt.target.name]: value
        })
    }

    return (
        <div className={"create"}>
            {success ? null : <h1>Not added</h1>}
            <h2>New Group</h2>
            <form className={"create-forms"} onSubmit={handleSubmit} onKeyDown={(e) => {
                e.key === 'Enter' && e.preventDefault();
            }}>
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
                <div className={"btn-group"}>
                    <button className={"btn back-btn"} onClick={() => navigate(-1)}>Back</button>
                    <button className={"btn send-btn"}>SEND</button>
                </div>
            </form>
        </div>
    )
}

export default CreateGroup
