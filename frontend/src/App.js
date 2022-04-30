import './App.css';
import {useEffect} from 'react';

function App() {
    const entries = new Map()
    entries.set(1, {
        url: 'http://homeassistant.local:8123',
        displayname: 'HomeAssistant',
        faviconLocation: '/static/icons/favicon.ico'
    })
    entries.set(2, {
        url: 'http://10.0.0.1',
        displayname: 'OpnSense',
        faviconLocation: '/ui/themes/opnsense/build/images/favicon.png'
    })
    entries.set(3, {
        url: 'http://10.0.0.2',
        displayname: 'DrayTek',
        faviconLocation: '/web/assets/img/favicon.png'
    })
    entries.set(4, {
        url: 'http://10.0.0.4',
        displayname: 'WiFi AP',
        faviconLocation: '/images/favicon.png'
    })
    entries.set(4, {
        url: 'http://10.0.0.93',
        displayname: 'Paperless NG',
        faviconLocation: '/favicon.ico'
    })
    entries.set(5, {
        url: 'http://10.0.0.91:32400/web',
        displayname: 'Plex',
        faviconLocation: '/favicon.ico'
    })
    entries.set(6, {
        url: 'https://10.0.0.8:8006',
        displayname: 'Proxmox',
        faviconLocation: '/pve2/images/logo-128.png'
    })
    entries.set(7, {
        url: 'http://10.0.0.5',
        displayname: 'Unraid',
        faviconLocation: '/webGui/images/green-on.png'
    })
    useEffect(() => {
        document.title = 'Bashdoard';
    }, []);

    return (
        <div className="bg-slate-900">
        <div className="container mx-auto h-screen">
            <div className="p-5 grid grid-cols-1 gap-2">
                {[...entries.entries()].map(([key, {url, faviconLocation, displayname}]) =>
                    <a key={key} href={url}>
                        <div className="bg-slate-800 flex items-center p-4">
                            <img className="rounded-lg border-2" alt="" width="64px" height="64px" src={url + faviconLocation}/>
                            <h1 className="ml-3 text-3xl text-white">{displayname}</h1>
                        </div>
                    </a>
                )}
            </div>
        </div>
        </div>
    );
}

export default App;
