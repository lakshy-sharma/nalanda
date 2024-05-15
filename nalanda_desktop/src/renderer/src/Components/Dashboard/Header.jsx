import Navigator from "./Navigation"
import nalandaLogo from '../../assets/logo.svg'

function Header({title="Nalanda", slogan="Your Personal Document Assistant"}) {
    return (
        <div className="container md:mx-auto shadow-md bg-slate-300 dark:bg-slate-700 mb-5 pt-5">
            <header className="container md:mx-auto text-center justify-center items-center">
                <h1 className="text-4xl font-medium flex text-center justify-center items-center">
                    <img src={nalandaLogo} width="40" height="40" className="mr-2"/>
                    {title}
                </h1>
                <p className="italic">
                    {slogan}
                    <br />
                </p>
            </header>
            <Navigator/>
        </div>
    )
}

export default Header