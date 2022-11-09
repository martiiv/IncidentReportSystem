import {BASE_URL} from "../constants/WebURL";

/**
 * Function to post data to an api
 *
 * @param url of the request we would like to send.
 * @param body post body the user is sending to the api
 * @returns {Promise<unknown>}
 */
export default function postModel(url, body) {
    return new Promise(function (resolve, reject) {
        const xhr = new XMLHttpRequest();
        xhr.open("POST", BASE_URL + url);
        xhr.setRequestHeader("Content-Type", "application/json");
        /*
            load event is also ok to use here,
            but readystatechange was giving me more descriptive errors
        */
        xhr.addEventListener('readystatechange', () => {
            if (xhr.readyState !== 4) {
                return;
            }
            if (xhr.status !== 201) {
                reject(new Error(JSON.stringify({
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
