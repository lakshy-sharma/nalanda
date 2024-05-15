import ContactForm from "./ContactForm";
import ContactInfo from "./ContactInfo";

function Footer() {
    return (
        <div className="container py-2 shadow-inner md:mx-auto text-center bg-slate-300 dark:bg-slate-700">
            <h2 className="text-center md:text-4xl font-bold">Contact Us</h2>    
            <div className="grid grid-cols-2">
                <section className="px-10">
                    <ContactInfo />
                </section>
                <section className="px-10">
                    <ContactForm />
                </section>
            </div>
            <footer className="mt-2">
                <p>&copy; 2024 Lakshy Sharma </p>
            </footer>
        </div>
    )
}

export default Footer