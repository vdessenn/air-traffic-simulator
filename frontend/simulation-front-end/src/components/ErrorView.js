import errorImg from "../assets/error-icon.svg";


export default function ErrorView() {
  return (
    <div>
    <div className="centered-div">
      <img src={errorImg} style={{width: 50, height: 50}} />
      <span>Nous ne parvenons pas à nous connecter au serveur. Veuillez vérifier :
      <ul>
        <li>que le serveur Go est bien lancé sur le port 7500</li>
        <li>que le format de données retourné est conforme au format attendu</li>
      </ul>
      </span>
    </div>
    </div>
    )
}