import './TagsInput.css';
import {useEffect, useState} from 'react'
import CreatableSelect from 'react-select/creatable';
import Select from "react-select";
import "../index.css"
import "../mainPages/Create.css"
import customStyles from "./SelectStyle";

/**
 * Code is mostly taken from https://dev.to/0shuvo0/lets-create-an-add-tags-input-with-react-js-d29 as this
 * example adds tags in a way we want to.
 * @returns {JSX.Element}
 * @constructor
 */
function TagsInput({setTagsFunc, data, setStateTestFunc}){
    const [options, setOptions] = useState([])

    /**
     * Function that will fill tags options when class is called upon.
     */
    useEffect(() => {
        if (typeof data === 'string'){
            setOptions([{value: 0, label: data}])
        }else{
            setOptions(data.map(item => ({value: item.id, label: item.name})))
        }
    }, [])

    /**
     * Function that will handle change when Select is selected.
     * @param e values of the selected option.
     */
    function handleChangeOptions(e) {
        if (setTagsFunc.length > 0) {
            setStateTestFunc(prevState => ({
                ...prevState,
                "countermeasure": data.filter(item => item.name === e.label).map(res => (res.description))[0],
                "tag": e.label,
            }))
        }
    }

    return (
        <div className={"tag-container"}>
            {typeof data === 'string' ?
                <Select styles={customStyles} defaultValue={{ label: data, value: 0 }} options={options} onChange={handleChangeOptions} className={"input-group select-group-log"}/>:
                <CreatableSelect styles={customStyles}  options={options} onChange={handleChangeOptions} className={"input-group select-group"}/>
            }
        </div>
    )
}

export default TagsInput
