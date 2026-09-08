import { TailSpin } from 'react-loader-spinner'

export default function LoadingView() {
  return (
    <div>
    <div className="centered-div">
      <TailSpin visible={true} color="black" />
      <span>Chargement des données de vol...</span>
    </div>
    </div>
    )
}