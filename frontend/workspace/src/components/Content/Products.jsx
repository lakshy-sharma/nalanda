function ProductsContent() {
    return(
        <div className="flex-center ml-11 mr-11 py-2">
            <header className="text-center md:text-4xl">
                Our Products and Services
                <hr className="w-48 h-1 mx-auto rounded md:my-5 bg-slate-600 dark:bg-slate-200"/>
            </header>

            <div>
                <h2 className="text-center md:text-2xl font-semibold">
                    Home Loan Processing Services
                </h2>
                <p className="md:text-lg mt-2 leading-relaxed">
                    We are authorized to help you in securing home loans from the following banks.
                    <ul className="list-disc ml-5">
                        <li>SBI Bank</li>
                        <li>HDFC Bank</li>
                        <li>ICICI Bank</li>
                    </ul>
                </p>
            </div>

            <div>
                <h2 className="text-center md:text-2xl font-semibold">
                    Mutual Funds Advisory Services
                </h2>
                <p className="md:text-lg mt-2 leading-relaxed">
                    We are certified to act as a mutual fund distributors by NISM and we are currrently affiliated with the following institutions to provide equity trading services.
                    <ul className="list-disc ml-5">
                        <li>Anandrathi</li>
                        <li>BSE/NSE</li>
                    </ul>
                </p>
            </div>

            <div>
                <h2 className="text-center md:text-2xl font-semibold">
                    Personalized Financial Services
                </h2>
                <p className="md:text-lg mt-2 leading-relaxed">
                    We provide the following one to one consultation services based on requirements of each individual.
                    <ul className="list-disc ml-5">
                        <li>Goal Based Personal Planning</li>
                        <li>Retirement Planning</li>
                        <li>Taxation and Refunds Management</li>
                    </ul>
                </p>
            </div>
        </div>
    )
}

export default ProductsContent