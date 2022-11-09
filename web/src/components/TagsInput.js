// src/components/TagsInput.js
import './TagsInput.css';
import {useEffect, useState} from 'react'
import CreatableSelect from 'react-select/creatable';
import Select from "react-select";

/**
 * TODO: Add API connection after Tags-table is created in the DB
 * TODO: Add write support to API when new tags are created
 * Code is mostly taken from https://dev.to/0shuvo0/lets-create-an-add-tags-input-with-react-js-d29 as this
 * example adds tags in a way we want to.
 * @returns {JSX.Element}
 * @constructor
 */

function TagsInput({setTagsFunc, data}){
    const [tags, setTags] = useState("")

    console.log(data)
    const [options, setOptions] = useState([])

    useEffect(() => {
        if (typeof data === 'string'){
            setOptions([{value: 0, label: data}])
        }else{
            setOptions(data.map(item => ({value: item.id, label: item.name})))
        }
    }, [tags])

    function handleChangeOptions(e) {
        setTagsFunc(e.label)
    }

    return (
        <div className={"tag-container"}>
            {typeof data === 'string' ?
                <Select isClearable defaultValue={{ label: data, value: 0 }} options={options} onChange={handleChangeOptions} className={"input-group"}/>:
                <CreatableSelect isClearable options={options} onChange={handleChangeOptions} className={"input-group"}/>

            }

        </div>
    )
}

export default TagsInput
