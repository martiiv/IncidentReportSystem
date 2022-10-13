import {Link} from "react-router-dom";
import {INCIDENT_URL} from "../constants/WebURL";


function Table(props){
    return(
        <div className={"table-div"} style={{display: "flex" , justifyContent: "center" }}>
            <table id={"table"}>
                <tr>
                    {props.row.map(item =>
                        <th key={item}>
                            {item}
                        </th>
                    )}
                </tr>

                {props.data.map(item => <tr>
                    <td>{item.date}</td>
                    <td>{item.receivingGroup}</td>
                    <td>{item.systemManager}</td>
                    <td>Tags</td>
                    <td>{item.countermeasure}</td>
                    <td>
                        <Link to={INCIDENT_URL +"/" + item.id}>
                            <button value={item.name}>
                                {item.name}
                            </button>
                        </Link>
                    </td>
                </tr>)
                }
            </table>
        </div>
    )
}

export default Table
