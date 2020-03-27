MongoDB, Go'yu destekleyen bir SDK'ya sahip. MongoDB ve Gorilla Toolkit kullanarak Go dilinde RESTful API geliştireceğiz. Yararlandığım videoya ek olarak biraz daha MVC mimarisine benzer bir yaklaşım denemeye çalıştım. Go'da yeni olduğum için pek başarabildim mi bilemiyorum. :)

Yararlandığım video: https://www.youtube.com/watch?v=oW7PMHEYiSk

Teşekkürler Nic Raboy!

BAĞIMLILIKLAR

MongoDB: MongoDB için Docker imajı kullanacağız.
    İmajı çekmek için: docker pull mongo:4.0.4
    İmajı çalıştırmak için: docker run -d -p 27017-27019:27017-27019 --name mongodb mongo:4.0.4

Gorilla Toolkit: RESTful yapı için kullanacağız. Yüklemek için gereken komut: go get github.com/gorilla/mux

MongoDB Go Driver: DB için gereken sürücü. Yüklemek için: go get go.mongodb.org/mongo-driver/mongo