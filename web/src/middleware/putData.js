import {BASE_URL} from "../constants/WebURL";

/**
 * Function to change data to the API
 *
 * @param url of the request we would like to send.
 * @param body put body the user is sending to the api
 * @returns {Promise<unknown>}
 */
export default function putModel(url, body) {
    return new Promise(function (resolve, reject) {
        const xhr = new XMLHttpRequest();
        xhr.open("PUT", BASE_URL + url);
        xhr.setRequestHeader("Content-Type", "application/json");
        /*
            load event is also ok to use here,
            but readystatechange was giving me more descriptive errors
        */
        xhr.addEventListener('readystatechange', () => {
            if (xhr.readyState !== 4) {
                return;
            }
            if (xhr.status !== 200) {
                reject((({
                    statusCode: xhr.status,
                    statusText: xhr.statusText,
                    text: xhr.responseText
                })));
            } else {
                resolve({
                        statusCode: xhr.status,
                        text: xhr.responseText
                    }
                );
            }
        });
        xhr.send(JSON.stringify(body));
    });
}
