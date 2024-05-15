import ImageSlider from "../Commons/ImageSlider"

function RecognitionContent() {
    const awards = [
        {"url":"./assets/sip_startup_winner.jpg", "description": "In 2017 we won the SIP startup award from NJ-eWealth."},
        {"url":"./assets/icici_pan_india_award.jpg", "description": "In 2021 we were recognised by the ICICI Home Loans for being among the top 10 Home loan disbursements all over India."},
    ]
    const certifications = [
        {"url":"./assets/strp_2009.jpg", "description": "In 2009 we were achieved the STRP certification from the NIIT."},
        {"url":"./assets/trp_2009.jpg", "description": "In 2009 we were achieved the TRP certification from the NIIT."},
        {"url":"./assets/nse_certificate.jpg", "description": "We are certified by the NSE and the BSE. This certificate highlights our commitment to compliance."},
        {"url":"./assets/bse_certificate.jpg", "description": "We are certified by the NSE and the BSE. This certificate highlights our commitment to compliance."},
        {"url":"./assets/nism_certificate.jpg", "description": "We have received the NISM certification from SEBI. This certificate highlights our capabilities in guiding you through the mutual fund markets."},
    ]

    const partnerships = [
        {"url":"./assets/sbi_asm_membership.jpg", "description": "We were granted the prestigious SBI ASM Club Membership for 2017-2018."},
        {"url":"./assets/anandrathi_partnership.jpg", "description": "We have been partners of the Anandrathi investment services since 5 years."},
        {"url":"./assets/prudent_partnership.jpg", "description": "We are partners of Prudent Connect and provide mutual fund services through Fundzbazar."},
    ]
    return(
        <div className="py-2">
            <section>
                <div className="flex-center ml-11 mr-11">
                    <h2 className="text-center text-4xl">Awards</h2>
                    <hr className="w-48 h-1 mx-auto rounded md:my-5 bg-slate-600 dark:bg-slate-200"/>
                    <div className="container flex items-center justify-center mb-16">
                        <ImageSlider slides={awards} sliderHeight={"740px"} sliderWidth={"1000px"}/>
                    </div>
                </div>
            </section>
            <section>
                <div className="flex-center ml-11 mr-11">
                    <h2 className="text-center text-4xl">Certifications</h2>
                    <hr className="w-48 h-1 mx-auto rounded md:my-5 bg-slate-600 dark:bg-slate-200"/>
                    <div className="container flex items-center justify-center mb-16">
                        <ImageSlider slides={certifications} sliderHeight={"740px"} sliderWidth={"1000px"}/>
                    </div>
                </div>
            </section>
            <section>
                <div className="flex-center ml-11 mr-11">
                    <h2 className="text-center text-4xl">Partnerships</h2>
                    <hr className="w-48 h-1 mx-auto rounded md:my-5 bg-slate-600 dark:bg-slate-200"/>
                    <div className="container flex items-center justify-center mb-16">
                        <ImageSlider slides={partnerships} sliderHeight={"740px"} sliderWidth={"1000px"}/>
                    </div>
                </div>
            </section>
        </div>
    )
}

export default RecognitionContent