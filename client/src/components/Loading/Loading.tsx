import "./loading.scss";

export default function IsFetching() {
  const loaderColor = {
    background: "#3498db",
  };

  return (
    <div className="loading-background">
      <div className="loading-bar">
        <div className="loading-circle-1" style={loaderColor} />
        <div className="loading-circle-2" style={loaderColor} />
      </div>
    </div>
  );
}
