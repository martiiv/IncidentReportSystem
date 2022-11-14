import {useEffect, useState} from "react";
import postModel from "../middleware/postData";
import {GROUPS_URL, RECEIVER_URL} from "../constants/WebURL";
import Select from 'react-select'
import "./Create.css"
import fetchData from "../middleware/FetchData";
import {useNavigate} from "react-router-dom";


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

    const handleSubmit = async (event) => {
        event.preventDefault();
        await postModel(RECEIVER_URL, warning)
            .then(() => navigate(-1))

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
        console.log(e.value)
        setWarning({
            ...warning, ["receiverGroup"]: e.value
        })
    }


    function handleChange(evt) {
        const value = evt.target.value;
        setWarning({
            ...warning, [evt.target.name]: value
        })
    }


    const customStyles = {
        control: (base, state) => ({
            ...base,
            background: "#023950",
            // Overwrittes the different states of border
            borderColor: state.isFocused ? "yellow" : "green",
            // Removes weird border around container
            boxShadow: state.isFocused ? null : null,
            "&:hover": {
                // Overwrittes the different states of border
                borderColor: state.isFocused ? "red" : "blue"
            }
        })
    };

    console.log(warning)
    return(
        <div className={"create"}>
            <h2>New Warning Receiver</h2>
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
                <button className={"btn send-btn"}>SEND</button>
            </form>
        </div>
    )
}
export default CreateWarningReceiver
