import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";

import "./error.scss";

import config from "../../config/config";

const CommonError: React.FC<any> = () => {
  const reduxDispatch = useDispatch();
  const navigate = useNavigate();

  return (
    <div className="errorContent">
      <div className="row">
        <div className="col-md-12">
          <div className="browser">
            <div className="controls">
              <i />
              <i />
              <i />
            </div>

            <div className="eye"></div>
            <div className="eye"></div>
            <div className="mouth">
              <div className="lips"></div>
              <div className="lips"></div>
              <div className="lips"></div>
              <div className="lips"></div>
              <div className="lips"></div>
              <div className="lips"></div>
              <div className="lips"></div>
              <div className="lips"></div>
            </div>
          </div>

          <div className="errorHeader">
            Unfortunately, something has gone wrong.
          </div>
          <div className="errorBody">
            We're unable to fulfill your request. Please
            <b
              style={{
                color: "#FFA500",
                cursor: "pointer",
              }}
              onClick={() => {
                window.location.reload();
              }}
            >
              {" "}
              refresh your browser
            </b>{" "}
            or{" "}
            <a
              style={{
                color: "#ff5e5b",
              }}
              onClick={() => {
                //TODO Add Clear Data Reducer
                navigate("/");
              }}
            >
              {" "}
              clear your site data
            </a>
            . If the error continues please contact our{" "}
            <b>
              <a
                style={{
                  color: "#4fd44c",
                }}
                href={`mailto:${config.supportEmail}?subject="${config.appName} Error"&body=Message`}
              >
                Support Team
              </a>
            </b>
            .
          </div>
        </div>
      </div>
    </div>
  );
};

export default CommonError;
