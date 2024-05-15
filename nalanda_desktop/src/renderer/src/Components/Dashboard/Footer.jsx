import { useState } from "react"
function Footer() {
    const [name, setName] = useState("")
    const [email, setEmail] = useState("")
    const [phoneNo, setPhoneNo] = useState("")
    const [subject, setSubject] = useState("")
    const [message, setMessage] = useState("")

    async function handleSubmit(event) {
        event.preventDefault();
        fetch(`https://script.google.com/macros/s/AKfycbw6QVNQsFi9cpxt2jFGQJsgltoD-0TzRhkGwPCLjPwJRcNGl0hdCOOzby8pPSzLwqY1SQ/exec?Name=${name}&Email=${email}&PhoneNo=${phoneNo}&Subject=${subject}&Message=${message}`, {
            method: "POST",
            cache: "no-cache",
            headers: {
              "Content-Type": "text/plain",
            },
        })
        .then((response) => response.json())
        .then((response) => {
          toast.success("Response Submitted");
        })
        .catch((response) => {
          toast.success("Submission Failed, Try Again Later.");
        });
    }

    return (
        <div className="container py-2 shadow-inner md:mx-auto text-center bg-slate-300 dark:bg-slate-700">
            <h2 className="text-center md:text-4xl font-bold">Contact Us</h2>    
            <div className="grid grid-cols-2">
                <section className="px-10">
                    <div>
                        <dl className="text-left">
                            <dt className="font-bold underline">Address</dt>
                            <dd className="font-light">
                                Shop No. 111, 1st Floor, Royal Dreams Shopping Complex, Vapi-Daman Road, Chala, Vapi, Gujarat, India. (396191)
                            </dd>
                            <dt className="mt-2 font-bold underline">Phone</dt>
                            <dd className="mt-2 text-m font-light">(+91) 9638349440</dd>
                            <dt className="mt-2 font-bold underline">Email</dt>
                            <dd className="text-m font-light"><a href="#">deepaktaxfin@gmail.com</a></dd>
                        </dl>
                    </div>
                </section>
                <section className="px-10">
                    <div className="flex items-center justify-center">
                        <div className="mx-auto w-full max-w-[750px]">
                            <form onSubmit={handleSubmit}>
                            <div className="grid grid-cols-2 mb-1">
                                <label htmlFor="name" className="block text-base font-medium">Full Name </label>
                                <input type="text"name="Name" id="name" placeholder="Full Name" className="w-full rounded-md border border-[#e0e0e0] bg-white py-1 px-1 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md" onChange={(e) => setName(e.target.value)}/>
                            </div>
                            <div className="grid grid-cols-2 mb-1">
                                <label htmlFor="email" className="mb-3 block text-base font-medium">Email Address</label>
                                <input type="email" name="Email" id="email" placeholder="Your Email (example@domain.com)" className="w-full rounded-md border border-[#e0e0e0] bg-white py-1 px-1 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md" onChange={(e) => setEmail(e.target.value)}/>
                            </div >
                            <div className="grid grid-cols-2 mb-1">
                                <label htmlFor="phone" className="mb-3 block text-base font-medium">Phone Number</label>
                                <input type="text" name="Phone" id="phone" placeholder="Your Phone Number" className="w-full rounded-md border border-[#e0e0e0] bg-white py-1 px-1 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md" onChange={(e) => setPhoneNo(e.target.value)}/>
                            </div >
                            <div className="grid grid-cols-2 mb-1">
                                <label htmlFor="subject" className="mb-3 block text-base font-medium" >Subject</label>
                                <input type="text" name="Subject" id="subject" placeholder="Subject of Discussion" className="w-full rounded-md border border-[#e0e0e0] bg-white py-1 px-1 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md" onChange={(e) => setSubject(e.target.value)}/>
                            </div>
                            <div className="grid grid-cols-2 mb-1">
                                <label htmlFor="message" className="mb-3 block text-base font-medium" >Message</label>
                                <textarea rows="4" name="Message" id="message" placeholder="Your Message" className="w-full resize-none rounded-md border border-[#e0e0e0] bg-white py-1 px-1 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md" onChange={(e) => setMessage(e.target.value)}/>
                            </div>
                            <div>
                                <input type="submit" className="hover:shadow-form rounded-md bg-white py-1 px-1 text-base font-semibold text-slate-600 hover:text-blue-800 outline-none" />
                            </div>
                            </form>
                        </div>
                    </div>
                </section>
            </div>
            <footer className="mt-2">
                <p>&copy; 2024 Lakshy Sharma </p>
            </footer>
        </div>
    )
}

export default Footer