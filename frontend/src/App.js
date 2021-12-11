import './App.css';
import { useState, useEffect } from 'react';

function App() {
  const [entries, setEntries] = useState(new Map());

  useEffect(() => {
    (async () => {
      let res = await fetch('http://localhost:8080/api/sync')
      let entries = await res.json()
      for (let entry of entries) {
        setEntries(prev => new Map(prev).set(entry.namespace + '-' + entry.name, entry))
      }
    })()
    let eventSource = new EventSource("http://localhost:8080/api/dashboardentries");
    eventSource.addEventListener('add', (evt) => {
      let entry = JSON.parse(evt.data)
      setEntries(prev => new Map(prev).set(entry.namespace + '-' + entry.name, entry))
    });
    eventSource.addEventListener('update', (evt) => {
      console.log('update')
      let entry = JSON.parse(evt.data)
      setEntries(prev => new Map(prev).set(entry.namespace + '-' + entry.name, entry))
    });
    eventSource.addEventListener('delete', (evt) => {
      let entry = JSON.parse(evt.data)
      setEntries(prev => {
        const newState = new Map(prev)
        newState.delete(entry.namespace + '-' + entry.name)
        return newState
      })
    });
  })
  return (
    <>
      <div className="container px-4 py-5" id="icon-grid">
        <h1 className="pb-2 border-bottom">Bashdoard</h1>
        <div className="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-4 py-5">
          {[...entries.entries()].map(([key, e]) =>
            <div key={key} className="col d-flex align-items-start">
              <img className="bi text-muted flex-shrink-0 me-3" alt="" width="28px" height="28px" src={e.url + e.faviconLocation}></img>
              <div>
                <h4 className="fw-bold mb-0"><a href={e.url}>{e.displayname}</a></h4>
              </div>
            </div>
          )}
        </div>
      </div>
    </>
  );
}

export default App;
