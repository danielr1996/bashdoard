import './App.css';
import {useEffect, useState} from 'react';

function App() {
    const [entries, setEntries] = useState(new Map());
    const [config, setConfig] = useState();
    const [configLoaded, setConfigLoaded] = useState(false)
    useEffect(() => {
        async function getConfig() {
            const res = await fetch(`${process.env.PUBLIC_URL}/config.json`)
            const config = await res.json()
            setConfig(config)
            setConfigLoaded(true)
        }
        getConfig()
    }, [])

    useEffect(() => {
        if (configLoaded) {
            async function getDashboardEntries() {
                let res = await fetch(`${config.api}/api/dashboardentries`)
                let entries = await res.json()
                for (let entry of entries) {
                    setEntries(prev => new Map(prev).set(entry.id, entry))
                }
            }
            getDashboardEntries()

            let eventSource = new EventSource(`${config.api}/api/dashboardentries/sse`);
            eventSource.addEventListener('add', (evt) => {
                let entry = JSON.parse(evt.data)
                setEntries(prev => new Map(prev).set(entry.id, entry))
            });
            eventSource.addEventListener('update', (evt) => {
                let entry = JSON.parse(evt.data)
                setEntries(prev => new Map(prev).set(entry.id, entry))
            });
            eventSource.addEventListener('delete', (evt) => {
                let entry = JSON.parse(evt.data)
                setEntries(prev => {
                    const newState = new Map(prev)
                    newState.delete(entry.id)
                    return newState
                })
            });
        }
    }, [config, configLoaded])

    useEffect(() => {
        document.title = 'Bashdoard';
    }, []);

    return (
        <div className="bg-slate-900">
        <div className="container mx-auto h-screen">
            <div className="p-5 grid grid-cols-1 gap-2">
                {[...entries.entries()].map(([key, {url, icon, name}]) =>
                    <a key={key} href={url}>
                        <div className="bg-slate-800 flex items-center p-4">
                            <img className="rounded-lg border-2" alt="" width="64px" height="64px" src={url + icon}/>
                            <h1 className="ml-3 text-3xl text-white">{name}</h1>
                        </div>
                    </a>
                )}
            </div>
        </div>
        </div>
    );
}

export default App;
