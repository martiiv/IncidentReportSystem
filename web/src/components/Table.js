import {Link} from "react-router-dom";
import {INCIDENT_URL} from "../constants/WebURL";

/**
 * Main table to display data for Warning receivers and Incident Log
 */
function Table({ onChangeValue, type, data, row }){
    return (
        <div className={"table-div"} style={{display: "flex", justifyContent: "center"}}>
            <table id={"table"}>
                <tbody>
                <tr>
                    {row.map(item =>
                        <th key={item}>
                            {item}
                        </th>
                    )}
                </tr>

                {data.map(item => <tr key={item.id}>

                    {type === "WR" &&
                        <td>
                            <label>
                                <input type="checkbox" value={item.id} onChange={onChangeValue}/>
                            </label>
                        </td>
                    }

                    <td>{type !== "WR" ? item.date : item.name}</td>
                    <td>{type !== "WR" ? item.receivingGroup : item.receiverGroup}</td>
                    <td>{type !== "WR"  ? item.name : item.company}</td>
                    <td>{type !== "WR" ? item.tag : item.receiverEmail}</td>
                    <td>{type !== "WR" ? item.countermeasure : item.phoneNumber}</td>
                    {type !== "WR" &&
                        <td key={item}>
                            <Link to={INCIDENT_URL + "/" + item.id}>
                                <button className={"btn tbl-btn"} value={item.name}>
                                    Manage Incident
                                </button>
                            </Link>
                        </td>
                    }
                </tr>)
                }
                </tbody>
            </table>
        </div>
    )}

export default Table
