/** @type {EventSource} */
let eventSource
let startTime
let id

function sendTime(task, endTime = new Date(), timeout = 2) {
    const elapsed = endTime - startTime

    fetch("/api/times", {
        method: "POST",
        headers: {
            "Content-Type": "application/type",
        },
        body: JSON.stringify({ id, ms: elapsed, task })
    }).then(() => {
        eventSource.close()
        window.location.href = "/wait?id=" + id
    }).catch(() => {
        setTimeout(() => sendTime(task, endTime, timeout * 2), timeout)
    })
}

window.history.pushState({}, document.title, window.location.pathname);

window.addEventListener("load", () => {
    console.log("loaded")

    id = parseInt(document.getElementById("UserID").innerText)

    if (!(id > 0)) {
        window.location.href = "/"
        return
    }

    startTime = new Date()

    eventSource = new EventSource("/wait/push");
    eventSource.addEventListener("message", e => {
        const eventData = e.data;
        console.log("Received data:", eventData);

        if (eventData)
            window.location.href = eventData + "?id=" + id
    });
})

