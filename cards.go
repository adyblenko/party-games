package main

import (
	"fmt"
	"io/ioutil"
    "net/http"
    "html/template"
    "regexp"
)

var playerAnswers = make(map[string]string)

type Page struct {
    Title string
    Body  []byte
}

type QA struct {
    Title string
    Body  string
}

var templates = template.Must(template.ParseFiles("browser.html", "mobile.html"))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

/*func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}*/

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

    //http.HandleFunc("/main/", indexHandler)
    http.HandleFunc("/mobile/", mobileHandler)
    http.HandleFunc("/browser/", browserHandler)
    http.HandleFunc("/save/", saveHandler)

    http.ListenAndServe(":8080", nil)
}

func mobileHandler(w http.ResponseWriter, r *http.Request) {   
   
    p:=&QA{Title: "State your shit", Body: ""}
    renderTemplate(w, "mobile", p)
}

func browserHandler(w http.ResponseWriter, r *http.Request) {   
    /*p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }*/
    p:=&QA{Title: "What can you say about this one?", Body: blackCards[1]}
    renderTemplate(w, "browser", p)
}

/*func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}*/

func saveHandler(w http.ResponseWriter, r *http.Request) {
    body := r.FormValue("body")
    player := r.URL.Path[len("/save/"):]
    playerAnswers[player] = body
    
    fmt.Fprintf(w, "PlayerAnswers %s!", playerAnswers)

    /*p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)*/
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *QA) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}

/*func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}*/

var blackCards = [...]string{
    "What ended my last relationship?",
    "__________. It's a trap!",
    "__________. High five, bro.",
    "I got 99 problems but __________ ain't one.",
    "How am I maintaining my relationship status?",
    "The new Chevy Tahoe. With the power and space to take __________ everywhere you go.",
    "When Pharaoh remained unmoved, Moses called down a Plague of __________.",
    "In the new Disney Channel Original Movie, Hannah Montana struggles with __________ for the first time. ",
    "While the United States raced the Soviet Union to the moon, the Mexican government funneled millions of pesos into research on __________.",
    "This is the way the world ends / This is the way the world ends / Not with a bang but with __________.",
    "Dear Abby, I'm having some trouble with __________ and would like your advice.",
    "__________. Betcha can't have just one!",
    "A romantic, candlelit dinner would be incomplete without __________.",
    "Coming to Broadway this season, __________: The Musical.",
    "What am I giving up for Lent?",
    "What will always get you laid?",
    "The Smithsonian Museum of Natural History has just opened an interactive exhibit on __________.",
    "What are my parents hiding from me?",
    "Maybe she's born with it. Maybe it's __________.",
    "Next from J.K. Rowling: Harry Potter and the Chamber of __________.",
    "What's that smell?",
    "Today on Maury: 'Help! My son is __________!'",
    "What gives me uncontrollable gas?",
    "__________. That's how I want to die.",
    "What's my anti-drug?",
    "War! What is it good for?",
    "What's George W. Bush thinking about right now?",
    "What don't you want to find in your Kung Pao chicken?",
    "I'm sorry, Professor, but I couldn't complete my homework because of __________.",
    "What would grandma find disturbing, yet oddly charming?",
    "What helps Obama unwind?",
    "A recent laboratory study shows that undergraduates have 50 percent less sex after being exposed to __________.",
    "What's that sound?",
    "Introducing Xtreme Baseball! It's like baseball, but with __________!",
    "What's my secret power?",
    "When I am President of the United States, I will create the Department of __________.",
    "After the earthquake, Sean Penn brought __________ to the people of Haiti.",
    "Next on ESPN2: The World Series of __________.",
    "What did the US airdrop to the children of Afghanistan?",
    "What's a girl's best friend?",
    "What's the new fad diet?",
    "What gets better with age?",
    "Why can't I sleep at night?",
    "What is Batman's guilty pleasure?",
    "Life for American Indians was forever changed when the White Man introduced them to __________.",
    "What's Teach for America using to inspire inner city students to succeed?",
    "I do not know with what weapons World War III will be fought, but World War IV will be fought with __________.",
    "White people like __________.",
    "What did Vin Diesel eat for dinner?",
    "In 1,000 years, when paper money is a distant memory, how will we pay for goods and services?",
    "Why am I sticky?",
    "__________: Kid-tested, mother-approved.",
    "Fun tip! When your man asks you to go down on him, try surprising him with __________ instead. ",
    "It's a pity that kids these days are all getting involved with __________.",
    "What's there a ton of in heaven?",
    "Alternative medicine is now embracing the curative powers of __________.",
    "What did I bring back from Mexico?",
    "Why do I hurt all over?",
    "What's the next Happy Meal toy?",
    "TSA guidelines now prohibit __________ on airplanes.",
    "Instead of coal, Santa now gives the bad children __________.",
    "What never fails to liven up the party?",
    "MTV's new reality show features eight washed-up celebrities living with __________.",
    "But before I kill you, Mr. Bond, I must show you __________.",
    "When I am a billionaire, I shall erect a 50-foot statue to commemorate __________.",
    "What will I bring back in time to convince people that I am a powerful wizard?",
    "During sex, I like to think about __________.",
    "In L.A. County Jail, word is you can trade 200 cigarettes for __________.",
    "The class field trip was completely ruined by __________.",
    "I get by with a little help from ______.",
    "Daddy, why is mommy crying?",
    "__________: Good to the last drop.",
    "I drink to forget __________.",
    "Here is the church / Here is the steeple / Open the doors / And there is __________.",
    "What's the most emo?",
    "How did I lose my virginity?",
    "The secret to a lasting marriage is communication, communication, and __________.",
    "My plan for world domination begins with __________.",
    "Next season on Man vs. Wild, Bear Grylls must survive in the depths of the Amazon with only __________ and his wits. ",
    "Science will never explain __________.",
    "The CIA now interrogates enemy agents by repeatedly subjecting them to __________.",
    "In Rome, there are whisperings that the Vatican has a secret room devoted to __________.",
    "When all else fails, I can always masturbate to __________.",
    "I learned the hard way that you can't cheer up a grieving friend with __________.",
    "In its new tourism campaign, Detroit proudly proclaims that it has finally eliminated __________.",
    "The socialist governments of Scandinavia have declared that access to __________ is a basic human right.",
    "In his new self-produced album, Kanye West raps over the sounds of __________.",
    "What's the gift that keeps on giving?",
    "When I pooped, what came out of my butt?",
    "In the distant future, historians will agree that __________ marked the beginning of America's decline.",
    "What has been making life difficult at the nudist colony?",
    "And I would have gotten away with it, too, if it hadn't been for __________.",
    "What brought the orgy to a grinding halt?",
    "A remarkable new study has shown that chimps have evolved their own primitive version of __________.",
    "Your persistence is admirable, my dear Prince. But you cannot win my heart with __________ alone.",
    "During his midlife crisis, my dad got really into __________.",
    "My new favorite porn star is Joey '__________' McGee.",
    "Before I run for president, I must destroy all evidence of my involvement with __________.",
    "This is your captain speaking. Fasten your seatbelts and prepare for __________. ",
    "In his newest and most difficult stunt, David Blaine must escape from __________.",
    "The Five Stages of Grief: denial, anger, bargaining, __________, acceptance.",
    "Members of New York's social elite are paying thousands of dollars just to experience __________.",
    "This month's Cosmo: 'Spice up your sex life by bringing __________ into the bedroom.'",
    "Little Miss Muffet Sat on a tuffet, Eating her curds and __________.",
    "My country, 'tis of thee, sweet land of __________.",
    "Next time on Dr. Phil: How to talk to your child about __________.",
    "Only two things in life are certain: death and __________.",
    "The healing process began when I joined a support group for victims of __________. ",
    "What's harshing my mellow, man?",
    "Charades was ruined for me forever when my mom had to act out __________.",
    "Tonight on 20/20: What you don't know about __________ could kill you.",
    "__________. Awesome in theory, kind of a mess in practice. ",
    "A successful job interview begins with a firm handshake and ends with __________.",
    "And what did YOU bring for show-and-tell?",
    "As part of his contract, Prince won't perform without __________ in his dressing room.",
    "As part of his daily regimen, Anderson Cooper sets aside 15 minutes for __________.",
    "Call the law offices of Goldstein and Goldstein, because no one should have to tolerate  __________ in the workplace.",
    "During high school, I never really fit in until I found __________ Club.",
    "Finally! A service that delivers __________ right to your door.",
    "Hey baby, come back to my place and I'll show you __________.",
    "I'm not like the rest of you. I'm too rich and busy for __________.",
    "In the seventh circle of Hell, sinners must endure __________ for all eternity. ",
    "Lovin' you is easy 'cause you're __________.",
    "Money can't buy me love, but it can buy me __________.",
    "My gym teacher got fired for adding __________ to the obstacle course.",
    "The blind date was going horribly until we discovered our shared interest in __________.",
    "To prepare for his upcoming role, Daniel Day-Lewis immersed himself in the world of __________.",
    "Turns out that __________-Man was neither the hero we needed nor wanted.",
    "What left this stain on my couch?",
    "The Japanese have developed a smaller, more efficient version of __________.",
    "I'm pretty sure I'm high right now, because I'm absolutely mesmerized by __________.",
    "What's fun until it gets weird?",
    "Man, this is bullshit. Fuck __________.",
    "2 AM in the city that never sleeps. The door swings open and *she* walks in, legs up to here. Something in her eyes tells me she's looking for __________.",
    "I'm sorry, sir, but we don't allow __________ at the country club. ",
    "It lurks in the night. It hungers for flesh. This summer, no one is safe from __________. ",
    "As King, how will I keep the peasants in line?  ",
    "Wes Anderson's new film tells the story of a precocious child coming to terms with __________.",
    "She's up all night for good fun. I'm up all night for __________.",
    "Alright, bros. Our frat house is condemned and all the hot slampieces are over at Gamma Phi. The time has come to commence Operation __________.",
    "How am I compensating for my tiny penis?",
    "Dear Leader Kim Jong Un, our village praises your infinite wisdom with a humble offering of __________.",
    "Do not fuck with me! I am literally __________ right now.",
    "This is the prime of my life. I'm young, hot, and full of __________. ",
    "You've seen the bearded lady! You've seen the ring of fire! Now, ladies and gentlemen, feast your eyes upon __________!",
    "Yo' mama so fat she __________!",
    "Hi, this is Jim from accounting. We noticed a $1,200 charge labeled '__________.' Can you explain?",
    "Having the worst day EVER. #__________",
    "Hi MTV! My name is Kendra, I live in Malibu, I'm into __________, and I love to have a good time. ",
    "In his farewell address, George Washington famously warned Americans about the dangers of __________.",
    "Life's pretty tough in the fast lane. That's why I never leave the house without __________.",
    "Don't forget! Beginning this week, Casual Friday will officially become '__________ Friday.'",
    "What's making things awkward in the sauna? ",
    "Now in bookstores: 'The Audacity of __________,' by Barack Obama. ",
    "Armani suit: $1,000. Dinner for two at that swanky restaurant: $300. The look on her face when you surprise her with __________: priceless.",
    "Do the Dew with our most extreme flavor yet! Get ready for Mountain Dew __________!",
    "Don't miss the action comedy of the year! One cop plays by the book. The other's only interested in one thing: __________.",
    "Here at the Academy for Gifted Children, we allow students to explore __________ at their own pace. ",
    "Why am I broke? ",
    "Well what do you have to say for yourself, Casey? This is the third time you've been sent to the principal's office for __________.",
    "And today's soup is Cream of __________. ",
    "I don't mean to brag, but they call me the Michael Jordan of __________. ",
    "Help me doctor, I've got __________ in my butt!",
    "In his new action comedy, Jackie Chan must fend off ninjas while also dealing with __________.",
    "WHOOO! God damn I love __________!",
    "What killed my boner?",
    "Do you lack energy? Does it sometimes feel like the whole world is __________? Zoloft.",
    "I work my ass off all day for this family, and this what I come home to? __________!?",
    "I have a strict policy. First date, dinner. Second date, kiss. Third date, __________.",
    "When I was a kid, we used to play Cowboys and __________.",
    "This is America. If you don't work hard, you don't succeed. I don't care if you're black, white, purple, or __________.",
    "You Won't Believe These 15 Hilarious __________ Bloopers!",
    "James is a lonely boy. But when he discovers a secret door in his attic, he meets a magical new friend: __________.",
    "If you had to describe me, the Card Czar, using only one of the cards in your hand, which one would it be?",
    "Don't worry, kid. It gets better. I've been living with __________ for 20 years. ",
    "My grandfather worked his way up from nothing. When he came to this country, all he had was the shoes on his feet and __________.",
    "Behind every powerful man is __________.",
    "You are not alone. Millions of Americans struggle with __________ every day. ",
    "Come to Dubai, where you can relax in our world-famous spas, experience the nightlife, or simply enjoy __________ by the poolside.",
    "'This is madness!' / 'No. THIS IS __________!'",
    "Listen, Gary, I like you. But if you want that corner office, you're going to have to show me __________.",
    "I went to the desert and ate of the peyote cactus. Turns out my spirit animal is __________.",
    "And would you like those buffalo wings mild, hot, or __________?",
    "The six things I could never do without: oxygen, facebook, chocolate, netflix, friends, and __________ LOL!",
    "Why won't you make love to me anymore? Is it __________?",
    "Puberty is a time of change. You might notice hair growing in new places. You might develop an interest in __________. This is normal.",
    "I'm sorry, Mrs. Chen, but there was nothing we could do. At 4:15 this morning, your son succumbed to __________.",
    "I'm Miss Tennessee, and if I could make the world better by changing one thing, I would get rid of __________.",
    "Tonight, we will have sex. And afterwards, if you'd like, a little bit of __________.",
    "Everybody join hands and close your eyes. Do you sense that? That's the presence of __________ in this room.",
    "To become a true Yanomamo warrior, you must prove that you can withstand __________ without crying out.",
    "Y'all ready to get this thing started? I'm Nick Cannon, and this is America's Got __________.",
    "Siskel and Ebert have panned __________ as 'poorly conceived' and 'sloppily executed.'",
    "Up next on Nickelodeon: 'Clarissa Explains __________.'",
    "I'm a bitch, I'm a lover, I'm a child, I'm __________.",
    "How did Stella get her groove back?",
    "Believe it or not, Jim Carrey can do a dead-on impression of __________.",
    "It's Morphin' Time! Mastodon! Pterodactyl! Triceratops! Sabertooth Tiger! __________!",
    "Tonight on SNICK: 'Are You Afraid of __________?'",
    "In what's being hailed as a major breakthrough, scientists have synthesized __________ in the lab.",
    "A study published in Nature this week found that __________ is good for you in small doses.",
    "What really killed the dinosaurs?",
    "Hey there, Young Scientists! Put on your labcoats and strap on your safety goggles, because today we're learning about __________!",
    "This season, Tim Allen must overcome his fear of __________ to save Christmas.",
    "After blacking out during New Year's Eve, I was awoken by __________.",
    "Every Christmas, my uncle gets drunk and tells the story about __________.",
    "Jesus is __________.",
    "On the third day of Christmas, my true love gave to me: three French hens, two turtle doves, and __________.",
    "What keeps me warm during the cold, cold winter?",
    "Wake up, America. Christmas is under attack by secular liberals and their __________.",
    "What's the one thing that makes an elf instantly ejaculate?",
    "I really hope my grandma doesn't ask me to explain __________ again. ",
    "Blessed are you, Lord our God, creator of the universe, who has granted us __________.",
    "Because they are forbidden from masturbating, Mormons channel their repressed sexual energy into __________.",
    "Kids these days with their iPods and their Internet. In my day, all we needed to pass the time was __________.",
    "GREETINGS HUMANS / I AM __________ BOT / EXECUTING PROGRAM",
    "Revealed: Why He Really Resigned! Pope Benedict's Secret Struggle with __________!",
    "Honey, Mommy and Daddy love you very much. But apparently Mommy loves __________ more than she loves Daddy.",
    "Dear Mom and Dad, Camp is fun. I like capture the flag. Yesterday, one of the older kids taught me about __________. I love you, Casey",
    "Why am I so tired?",
    "Behold the Four Horsemen of the Apocalypse! War, Famine, Death, and __________. ",
    "Looking to earn some big bucks? Learn how to make __________ work for you!",
    "How are the writers of Cards Against Humanity spending your $25?",
    "I'm not going to lie. I despise __________. There, I said it.",
    "A wise man said, 'Everything is about sex. Except sex. Sex is about __________.'",
    "Corruption. Betrayal. __________. Coming soon to Netflix, 'House of __________.'",
    "Cancel all my meetings. We've got a situation with __________ that requires my immediate attention.",
    "Our relationship is strictly professional. Let's not complicate things with __________.",
    "Meryl Streep + __________ = Oscar.",
    "Hey Reddit! I'm __________. Ask me anything.",
    "I'm going on a cleanse this week. Nothing but kale juice and __________.",
    "Well, gentlemen! If you'll excuse me, I have a date with __________.",
    "Our demands are quite simple. We will execute a hostage every hour until you bring us _______."}

var blackCards2 = [...]string{
    "I never truly understood __________ until I encountered __________.",
    "When I was tripping on acid, __________ turned into __________.",
    "Introducing the amazing superhero/sidekick duo! It's __________ and __________!",
    "In a world ravaged by __________, our only solace is __________.",
    "They said we were crazy. They said we couldn't put __________ inside of __________. They were wrong.",
    "That's right, I killed __________. How, you ask? __________.",
    "__________ is a slippery slope that leads to __________.",
    "In M. Night Shyamalan's new movie, Bruce Willis discovers that __________ had really been __________ all along.",
    "And the Academy Award for __________ goes to __________.",
    "For my next trick, I will pull __________ out of __________.",
    "Lifetime presents '__________, the Story of __________.'",
    "Step 1: __________. Step 2:__________. Step 3: Profit.",
    "Dear Sir or Madam, We regret to inform you that the Office of __________ has denied your request for __________.",
    "In a pinch, __________ can be a suitable substitute for __________.",
    "Michael Bay's new three-hour action epic pits __________ against __________.",
    "You haven't truly lived until you've experienced __________ and __________ at the same time.",
    "__________ would be woefully incomplete without __________.",
    "My mom freaked out when she looked at my browser history and found __________.com/__________.",
    "If God didn't want us to enjoy __________, he wouldn't have given us __________.",
    "I spent my whole life working towards __________, only to have it ruined by __________.",
    "Before __________, all we had was __________.",
    "__________: Hours of fun. Easy to use. Perfect for __________!",
    "After months of practice with __________, I think I'm finally ready for __________.",
    "Having problems with __________? Try __________!",
    "Listen, son. If you want to get involved with __________, I won't stop you. Just steer clear of __________.",
    "My life is ruled by a vicious cycle of __________ and __________.",
    "When you get right down to it, __________ is just __________.",
    "With enough time and pressure, __________ will turn into __________.",
    "__________ may pass, but __________ will last forever.",
    "In return for my soul, the Devil promised me __________, but all I got was __________.",
    "We never did find __________, but along the way, we sure learned a lot about __________.",
    "Forget everything you know about __________, because now we've supercharged it with __________!",
    "I am become __________, destroyer of __________.",
    "You guys, I saw this crazy movie last night. It opens on __________, and then there's some stuff about __________, and then it ends with __________. (DRAW 2, PICK 3)",
    "If you can't handle __________, you'd better stay away from __________.",
    "Honey, I have a new roleplay I want to try tonight!  You can be __________, and I'll be __________.",
    "Adventure. Romance. __________. From Paramount Pictures, '__________.'",
    "This year's hottest album is '__________' by __________.",
    "Oprah's book of the month is '__________ For __________: A story of hope.'",
    "In the beginning, there was __________. And the Lord said 'Let there be __________.'",
    "__________ will never be the same after __________.",
    "Every step towards __________ gets me a little bit closer to __________. ",
    "Patient presents with __________. Likely a result of __________.",
    "Well if __________ is good enough for __________, it's good enough for me.",
    "Heed my voice, mortals! I am the god of __________, and I will not tolerate __________!",
    "In line with our predictions, we find a robust correlation between __________ and __________ (p < .05).",
    "In an attempt to recreate conditions just after the Big Bang, physicists at the LHC are observing collisions between __________ and __________.",
    "Today on Mythbusters, we find out how long __________ can withstand __________.",
    "But wait, there's more! If you order __________ in the next 15 minutes, we'll throw in __________ absolutely free!",
    "Here's what you can expect for the new year. Out: __________. In: __________.",
    "Today on Buzzfeed: 10 pictures of __________ that look like __________!",
    "A curse upon thee! Many years from now, just when you think you're safe, __________ shall turn into __________.",
    "Because you enjoyed __________, we thought you'd like __________."}

