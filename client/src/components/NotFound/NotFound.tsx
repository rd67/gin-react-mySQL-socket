import { Link } from "react-router-dom";

import "./notFound.scss";

export default function NotFound() {
  return (
    <div id="notfound">
      <div className="notfound">
        <div className="notfound-404">
          <h1>
            4<span>0</span>4
          </h1>
        </div>
        <br />
        <br />
        <br />
        <p>
          The page you are looking for might have been removed had its name
          changed or is temporarily unavailable.
        </p>
        <Link to="/">Home Page</Link>
      </div>
    </div>
  );
}
