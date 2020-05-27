export default async config => {
    const res = await fetch("/setup/", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(config),
    });
    if (res.status !== 204) throw new Error(await res.text());
};
