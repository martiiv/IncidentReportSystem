// src/components/TagsInput.js
import './TagsInput.css';
import {useEffect, useState} from 'react'
import dummyData from "../constants/DummyData";

/**
 * Code is mostly taken from https://dev.to/0shuvo0/lets-create-an-add-tags-input-with-react-js-d29 as this
 * example adds tags in a way we want to.
 * @returns {JSX.Element}
 * @constructor
 */
function TagsInput({setTagsFunc, data, changeable}){
    const test = data
    const [tags, setTags] = useState([test])
    console.log(test);

    function handleKeyDown(e){
        if(e.key !== 'Enter') return
        const value = e.target.value
        if(!value.trim()) return
        setTags([...tags, value])
        e.target.value = ''
    }

    function removeTag(index){
        setTags(tags.filter((el, i) => i !== index))
    }

    useEffect(() => {
        setTagsFunc(tags);
    });

    return (
        <div className={"tag-container"}>
            {tags.map((tag, index) => (
                tag !== undefined &&
                <div className="tag-item" key={index}>
                    <span className="text-output">{tag}</span>
                    {
                        changeable ?  <span className="close" onClick={() => removeTag(index)}>&times;</span> : null
                    }
                </div>
            ))
            }
            {changeable ?
                <input onKeyDown={handleKeyDown} type="text" className={"tags-input"} placeholder="Type somthing" /> : null
            }
        </div>
    )
}

export default TagsInput
